package models

import (
	"encoding/hex"
	"fmt"
	"math/big"

	constant "fclink.cn/grabblock/common"
	"fclink.cn/grabblock/entity"
	"fclink.cn/grabblock/tools"
	"fclink.cn/grabblock/tools/log"
	"github.com/golang/protobuf/proto"
)

// trans
type MainTransaction struct {
	TransactionId       string `gorm:"column:transaction_id;type:varchar(100);primary_key;"`
	TransactionHash     string `gorm:"column:transaction_hash;type:varchar(100);"`
	TransactionIndex    int    `gorm:"column:transaction_index"`
	TransactionOutCount int    `gorm:"column:transaction_out_count"`
	BlockHeight         int64  `gorm:"column:block_height;type:bigint"`
	BlockHash           string `gorm:"column:block_hash;type:varchar(100);"`
	PeerId              string `gorm:"column:peer_id;type:varchar(100);"`
	Nonce               int32  `gorm:"column:nonce"`
	FeeHi               int64  `gorm:"column:fee_hi"`
	FeeLow              int64  `gorm:"column:fee_low"`
	InnerCodetype       int32  `gorm:"column:inner_code"`     //交易类型 //0普通账户,1=多重签名账户，2=20合约账户，3=721合约账户,4=CVM合约,5=JSVM合约(可并行)
	SubCodeType         int32  `gorm:"column:sub_inner_code"` // 2=创建token 3=token转账 20=创建合约 21=合约调用
	CodeData            string `gorm:"column:code_data;type:varchar(1000)"`
	ExtData             string `gorm:"column:ext_data;type:varchar(2000)"`
	FromAddress         string `gorm:"column:from_address;type:varchar(100)"`
	ToAddress           string `gorm:"column:to_address;type:varchar(100)"`   //合约地址或者token地址，
	CtAddress           string `gorm:"column:ct_address;type:varchar(100)"`   //token out address
	Amount              string `gorm:"column:amount;type:varchar(100)"`       //orm big.Int create 失效
	TotalAmount         string `gorm:"column:total_amount;type:varchar(100)"` //orm big.Int create 失效
	Status              string `gorm:"column:status;type:varchar(10)"`
	StatusResult        string `gorm:"column:status_result;type:varchar(100)"`
	StatusHash          string `gorm:"column:status_hash;type:varchar(100)"`
	StatusTime          int64  `gorm:"column:status_time;"`
	Timestamp           int64  `gorm:"column:timestamp;"`
	AcceptTime          int64  `gorm:"column:accept_time;"`
	NodeAddress         string `gorm:"column:node_address;type:varchar(100)"`
	Isdone              int32  `gorm:"column:isdone;"`                  //是否解析合约 0=待解析 1=解析完成 2=不需要解析
	FailCount           int32  `gorm:"column:fail_count;"`              //解析错误次数
	SysId               string `gorm:"column:sys_id;type:varchar(100)"` //系统ID
	BizId               string `gorm:"column:biz_id;type:varchar(100)"` //业务域ID
	Model
}

/**
  统计合约调用次数
*/
func (m *MainTransaction) CountToAddress(toAddress string) int64 {
	count := int64(0)
	db.Model(MainTransaction{}).Where("to_address=?", toAddress).Count(&count)
	return count
}

func GetMainTransByContract(size int64) []MainTransaction {
	data := make([]MainTransaction, 0)
	db.Order("block_height asc,Timestamp asc").Where("isdone=0 and inner_code=4 and status=? and to_address is not null", constant.TRANSACTION_SUCCESS).Limit(size).Find(&data)
	return data
}
func GetMainTransByFile(size int64) []MainTransaction {
	data := make([]MainTransaction, 0)
	db.Order("block_height asc,Timestamp asc").Where("isdone=0 and inner_code=6 and status=?", constant.TRANSACTION_SUCCESS).Limit(size).Find(&data)
	return data
}
func UpdateIsdone(TransactionId string, isdone int64) error {
	return db.Model(&MainTransaction{
		TransactionId: TransactionId,
	}).UpdateColumn("isdone", isdone).Error
}

//insert transaction 0x01 0xff all insert
func AddMainTransaction(m *entity.TransactionInfo) ([]*MainTransaction, error) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		log.Errorf("err:%s", err.Error())
		return nil, err
	}
	accountTrans := make([]*MainTransaction, 0)
	t := &MainTransaction{
		TransactionId:    fmt.Sprintf("%s_%d", m.Hash, 0),
		TransactionHash:  m.Hash,
		TransactionIndex: 0,
		BlockHeight:      m.Status.Height,
		BlockHash:        m.Status.Hash,
		Nonce:            m.Body.Nonce,
		FeeHi:            m.Body.FeeHi,
		FeeLow:           m.Body.FeeLow,
		InnerCodetype:    m.Body.InnerCodetype,
		CodeData:         m.Body.CodeData,
		ExtData:          m.Body.ExtData,
		FromAddress:      m.Body.Address,
		Status:           m.Status.Status,
		StatusResult:     m.Status.Result,
		StatusTime:       m.Status.Timestamp / 1000,
		Timestamp:        m.Body.Timestamp / 1000,
		AcceptTime:       m.Accepttimestamp / 1000,
		SysId:            m.Body.SysId,
		BizId:            m.Body.BizCode,
	}
	if m.Node != nil {
		t.PeerId = m.Node.Nid
		t.NodeAddress = m.Node.Address
	}
	//交易总金额
	totalAmount := big.NewInt(0)
	//rc20 code_data解析
	rc20 := &entity.ContractRC20{}
	if m.Body.InnerCodetype == constant.INNER_CODE_TYPE_RC20 {
		v, _ := hex.DecodeString(m.Body.CodeData[2:])
		if err := proto.Unmarshal(v, rc20); err != nil {
			tx.Rollback()
			log.Errorf("err:%s", err.Error())
			return nil, err
		}
		for _, item := range rc20.Values {
			totalAmount.Add(totalAmount, tools.HexToBigInt(hex.EncodeToString(item)))
		}
	}

	if len(m.Body.Outputs) == 0 {
		switch codetype := m.Body.InnerCodetype; codetype {
		case constant.INNER_CODE_TYPE_CVM:
			t.SubCodeType = constant.INNER_CODE_TYPE_CVM_CREATE
			t.Amount = totalAmount.String()
			t.TotalAmount = totalAmount.String()
			if m.Status.Status == constant.TRANSACTION_SUCCESS {
				t.ToAddress = m.Status.Result
			}
			if err := tx.Create(t).Error; err != nil {
				tx.Rollback()
				log.Errorf("err:%s", err.Error())
				return nil, err
			}
			accountTrans = append(accountTrans, t)
		case constant.INNER_CODE_TYPE_RC20:
			t.SubCodeType = constant.INNER_CODE_TYPE_RC20_CREATE
			t.Amount = totalAmount.String()
			t.TotalAmount = totalAmount.String()
			if m.Status.Status == constant.TRANSACTION_SUCCESS {
				t.ToAddress = m.Status.Result
			}
			if len(rc20.Tos) == len(rc20.Values) {
				for index, item := range rc20.Tos {
					t.TransactionId = fmt.Sprintf("%s_%d", m.Hash, index)
					t.Amount = tools.HexToBigInt(hex.EncodeToString(rc20.Values[index])).String()
					t.CtAddress = hex.EncodeToString(item)
					t.TransactionIndex = index
					t.TransactionOutCount = len(rc20.Tos)
					if err := tx.Create(t).Error; err != nil {
						tx.Rollback()
						log.Errorf("err:%s", err.Error())
						return nil, err
					}
					accountTrans = append(accountTrans, t)
				}
			}
		default:
			t.Amount = totalAmount.String()
			t.TotalAmount = totalAmount.String()
			if err := tx.Create(t).Error; err != nil {
				tx.Rollback()
				log.Errorf("err:%s", err.Error())
				return nil, err
			}
			accountTrans = append(accountTrans, t)
		}
	} else {
		switch codetype := m.Body.InnerCodetype; codetype {
		case constant.INNER_CODE_TYPE_CVM:
			t.ToAddress = m.Body.Outputs[0].Address
			t.Amount = totalAmount.String()
			t.TransactionOutCount = len(m.Body.Outputs)
			t.TotalAmount = totalAmount.String()
			t.SubCodeType = constant.INNER_CODE_TYPE_CVM_CALL
			if err := tx.Create(t).Error; err != nil {
				tx.Rollback()
				log.Errorf("err:%s", err.Error())
				return nil, err
			}
			accountTrans = append(accountTrans, t)
		case constant.INNER_CODE_TYPE_RC20:
			t.Amount = tools.HexToBigInt(m.Body.Outputs[0].Amount).String()
			t.TransactionOutCount = len(m.Body.Outputs)
			t.TotalAmount = totalAmount.String()
			t.SubCodeType = constant.INNER_CODE_TYPE_RC20_TRANSFER
			t.ToAddress = m.Body.Outputs[0].Address
			if len(rc20.Tos) == len(rc20.Values) {
				for index, item := range rc20.Tos {
					t.TransactionId = fmt.Sprintf("%s_%d", m.Hash, index)
					t.Amount = tools.HexToBigInt(hex.EncodeToString(rc20.Values[index])).String()
					t.CtAddress = hex.EncodeToString(item)
					t.TransactionIndex = index
					t.TransactionOutCount = len(rc20.Tos)
					if err := tx.Create(t).Error; err != nil {
						tx.Rollback()
						log.Errorf("err:%s", err.Error())
						return nil, err
					}
					accountTrans = append(accountTrans, t)
				}
			}
		default:
			for _, trans := range m.Body.Outputs {
				totalAmount.Add(totalAmount, tools.HexToBigInt(trans.Amount))
			}
			for i, trans := range m.Body.Outputs {
				t.ToAddress = trans.Address
				t.Amount = tools.HexToBigInt(trans.Amount).String()
				t.TransactionId = fmt.Sprintf("%s_%d", m.Hash, i)
				t.TransactionIndex = i
				t.TransactionOutCount = len(m.Body.Outputs)
				t.TotalAmount = totalAmount.String()
				if err := tx.Create(t).Error; err != nil {
					tx.Rollback()
					log.Errorf("err:%s", err.Error())
					return nil, err
				}
				//引用类型必须深度copy生成新的对象
				t2 := new(MainTransaction)
				if err := tools.Copy(t2, t); err != nil {
					panic(err.Error())
				}
				accountTrans = append(accountTrans, t2)
			}
		}
	}

	//main_account_transaction
	sql := AddMainAccountTransactionGroup(accountTrans)
	if err := tx.Exec(sql).Error; err != nil {
		tx.Rollback()
		log.Errorf("err:%s", err.Error())
		return nil, err
	}

	if m.Body.InnerCodetype == constant.INNER_CODE_TYPE_CVM {
		//发合约outputs=0
		if len(m.Body.Outputs) == 0 && m.Status.Status == constant.TRANSACTION_SUCCESS {
			contractCount := 0
			tx.Model(&MainContract{}).Where("contract_address=?", m.Status.Result).Count(&contractCount)
			if contractCount == 0 {
				if err := tx.Create(&MainContract{
					ContractAddress:  m.Status.Result,
					ContractBytecode: m.Body.CodeData,
					CreatorTime:      m.Accepttimestamp / 1000,
					CreatorAddress:   m.Body.Address,
					BlockHeight:      m.Status.Height,
				}).Error; err != nil {
					tx.Rollback()
					log.Errorf("err:%s", err.Error())
					return nil, err
				}
			}
		}
	}
	//main_token
	if m.Body.InnerCodetype == constant.INNER_CODE_TYPE_RC20 {
		//发token outputs=0
		if len(m.Body.Outputs) == 0 {
			if m.Status.Status == constant.TRANSACTION_SUCCESS {
				tokenCount := 0
				tx.Model(&MainToken{}).Where("address=?", m.Status.Result).Count(&tokenCount)
				if tokenCount == 0 {
					if err := tx.Create(&MainToken{
						Symbol:         rc20.Symbol,
						Name:           rc20.Name,
						Decimals:       rc20.Decimals,
						Address:        m.Status.Result,
						BlockHeight:    m.Status.Height,
						CreatorAddress: m.Body.Address,
						CreatorTime:    m.Accepttimestamp / 1000,
						Nonce:          m.Body.Nonce,
						TotalSupply:    totalAmount.String(),
					}).Error; err != nil {
						tx.Rollback()
						log.Errorf("err:%s", err.Error())
						return nil, err
					}
				}
			}
		}
	}
	//service.PushTrans(accountTrans)
	return accountTrans, tx.Commit().Error
}
