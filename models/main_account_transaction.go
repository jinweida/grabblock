package models

import (
	"bytes"
	"fmt"
	"math/big"
	"time"

	"fclink.cn/grabblock/common"
	"fclink.cn/grabblock/entity"
	"fclink.cn/grabblock/tools"
)

// account_transaction
type MainAccountTransaction struct {
	Id                  int64    `gorm:"column:id;type:bigint;AUTO_INCREMENT;primary_key"`
	Address             string   `gorm:"column:address;type:varchar(100)"`
	TransactionHash     string   `gorm:"column:transaction_hash;type:varchar(100)"`
	TransactionIndex    int      `gorm:"column:transaction_index"`
	TransactionType     int      `gorm:"column:transaction_type;type:varchar(100)"`
	FromAddress         string   `gorm:"column:from_address;type:varchar(100)"`
	Toddress            string   `gorm:"column:to_address;type:varchar(100)"`
	Direction           string   `gorm:"column:direction;type:varchar(5)"`
	Amount              *big.Int `gorm:"column:amount;type:DECIMAL(32,0)"`
	TotalAmount         *big.Int `gorm:"column:total_amount;type:DECIMAL(32,0)"`
	BlockHeight         int64    `gorm:"column:block_height;type:bigint"`
	BlockHash           string   `gorm:"column:block_hash;type:varchar(100);"`
	TransactionOutCount int      `gorm:"column:transaction_out_count"`
	TokenSymbol         string   `gorm:"column:token_symbol;type:varchar(100)"`
	FromTag             string   `gorm:"column:from_tag;type:varchar(50)"`
	ToTag               string   `gorm:"column:to_tag;type:varchar(50)"`
	Status              string   `gorm:"column:status;type:varchar(10)"`
	StatusResult        string   `gorm:"column:status_result;type:varchar(500)"`
	StatusHash          string   `gorm:"column:status_hash;type:varchar(100)"`
	StatusTime          int64    `gorm:"column:status_time;"`
	Timestamp           int64    `gorm:"column:timestamp;"`
	AcceptTime          int64    `gorm:"column:accept_time;"`
	SysId               string   `gorm:"column:sys_id;type:varchar(100)"` //系统ID
	BizId               string   `gorm:"column:biz_id;type:varchar(100)"` //业务域ID
	Model
}

func AddMainAccountTransactionGroup(list []*MainTransaction) string {
	timeTemplate := "2006-01-02 15:04:00"
	var buffer bytes.Buffer
	buffer.WriteString("INSERT INTO `main_account_transaction` (`address`,`transaction_hash`,`transaction_index`,`transaction_type`,")
	buffer.WriteString("`from_address`,`to_address`,`direction`,`amount`,`block_height`,")
	buffer.WriteString("`status`,`status_hash`,`timestamp`,`accept_time`,`created`,`updated`,`total_amount`," +
		"`transaction_out_count`,`inner_code`,`sub_inner_code`,`ct_address`,`block_hash`,`sys_id`,`biz_id`) VALUES ")

	for i, trans := range list {
		if i > 0 {
			buffer.WriteString(",")
		}
		//from
		buffer.WriteString(fmt.Sprintf("('%s','%s',%d,%d,'%s','%s','%s',%s,%d,'%s','%s','%d','%d','%s','%s','%s','%d','%d','%d','%s','%s','%s','%s')",
			trans.FromAddress, trans.TransactionHash, trans.TransactionIndex,
			0, trans.FromAddress, trans.ToAddress, common.TRANSACTION_OUT_TEXT,
			trans.Amount, trans.BlockHeight, trans.Status, trans.StatusHash,
			trans.Timestamp,
			trans.AcceptTime,
			time.Now().Format(timeTemplate),
			time.Now().Format(timeTemplate),
			trans.TotalAmount,
			trans.TransactionOutCount,
			trans.InnerCodetype,
			trans.SubCodeType,
			trans.CtAddress,
			trans.BlockHash,
			trans.SysId,
			trans.BizId))
		//to
		if trans.SubCodeType == common.INNER_CODE_TYPE_RC20_TRANSFER {
			buffer.WriteString(fmt.Sprintf(",('%s','%s',%d,%d,'%s','%s','%s',%s,%d,'%s','%s','%d','%d','%s','%s','%s','%d','%d','%d','%s','%s','%s','%s')",
				trans.CtAddress, trans.TransactionHash, trans.TransactionIndex,
				0, trans.FromAddress, trans.ToAddress, common.TRANSACTION_IN_TEXT,
				trans.Amount, trans.BlockHeight, trans.Status, trans.StatusHash,
				trans.Timestamp,
				trans.AcceptTime,
				time.Now().Format(timeTemplate),
				time.Now().Format(timeTemplate),
				trans.TotalAmount,
				trans.TransactionOutCount,
				trans.InnerCodetype,
				trans.SubCodeType,
				trans.CtAddress,
				trans.BlockHash,
				trans.SysId,
				trans.BizId))
		} else {
			if len(trans.ToAddress) > 0 {
				buffer.WriteString(fmt.Sprintf(",('%s','%s',%d,%d,'%s','%s','%s',%s,%d,'%s','%s','%d','%d','%s','%s','%s','%d','%d','%d','%s','%s','%s','%s')",
					trans.ToAddress, trans.TransactionHash, trans.TransactionIndex,
					0, trans.FromAddress, trans.ToAddress, common.TRANSACTION_IN_TEXT,
					trans.Amount, trans.BlockHeight, trans.Status, trans.StatusHash,
					trans.Timestamp,
					trans.AcceptTime,
					time.Now().Format(timeTemplate),
					time.Now().Format(timeTemplate),
					trans.TotalAmount,
					trans.TransactionOutCount,
					trans.InnerCodetype,
					trans.SubCodeType,
					trans.CtAddress,
					trans.BlockHash,
					trans.SysId,
					trans.BizId))
			}
		}

	}
	buffer.WriteString(";")
	return buffer.String()
}

func AddMainAccountTransaction(m *entity.TransactionInfo, totalAmount *big.Int) string {
	timeTemplate := "2006-01-02 15:04:00"
	var buffer bytes.Buffer
	buffer.WriteString("INSERT INTO `main_account_transaction` (`address`,`transaction_hash`,`transaction_index`,`transaction_type`,")
	buffer.WriteString("`from_address`,`to_address`,`direction`,`amount`,`block_height`,")
	buffer.WriteString("`status`,`status_hash`,`timestamp`,`accept_time`,`created`,`updated`,`total_amount`,`transaction_out_count`) VALUES ")
	if len(m.Body.Outputs) == 0 {
		buffer.WriteString(fmt.Sprintf("('%s','%s',%d,%d,'%s','%s','%s',%d,%d,'%s','%s','%d','%d','%s','%s','%d','%d')",
			m.Body.Address, m.Hash, 0, m.Body.InnerCodetype, m.Body.Address, "", common.TRANSACTION_OUT_TEXT,
			0, m.Status.Height, m.Status.Status, m.Status.Hash,
			m.Status.Timestamp,
			m.Accepttimestamp/1000,
			time.Now().Format(timeTemplate),
			time.Now().Format(timeTemplate), 0, 0))
	}

	for index, trans := range m.Body.Outputs {
		//from
		buffer.WriteString(fmt.Sprintf("('%s','%s',%d,%d,'%s','%s','%s',%d,%d,'%s','%s','%d','%d','%s','%s','%d','%d')",
			m.Body.Address, m.Hash, index, m.Body.InnerCodetype, m.Body.Address, trans.Address, common.TRANSACTION_IN_TEXT,
			tools.HexToBigInt(trans.Amount), m.Status.Height, m.Status.Status, m.Status.Hash,
			m.Status.Timestamp/1000,
			m.Accepttimestamp/1000,
			time.Now().Format(timeTemplate),
			time.Now().Format(timeTemplate),
			totalAmount,
			len(m.Body.Outputs)))
		//to
		buffer.WriteString(fmt.Sprintf(",('%s','%s',%d,%d,'%s','%s','%s',%d,%d,'%s','%s','%d','%d','%s','%s','%d','%d')",
			trans.Address, m.Hash, index, m.Body.InnerCodetype, m.Body.Address, trans.Address, common.TRANSACTION_IN_TEXT,
			tools.HexToBigInt(trans.Amount), m.Status.Height, m.Status.Status, m.Status.Hash,
			m.Status.Timestamp/1000,
			m.Accepttimestamp/1000,
			time.Now().Format(timeTemplate),
			time.Now().Format(timeTemplate),
			totalAmount,
			len(m.Body.Outputs)))

	}
	buffer.WriteString(";")
	return buffer.String()
}

//统计biz上链人数
func (m *MainAccountTransaction) CountOnlineUser(bizid string) int64 {
	count := int64(0)
	db.Model(m).Select("DISTINCT(address)").Where("biz_id=?", bizid).Count(&count)
	return count
}

//统计biz下系统上链人数
func (m *MainAccountTransaction) CountOnlineUserInsys(bizid string, sysid string) int64 {
	count := int64(0)
	db.Model(m).Select("DISTINCT(address)").Where("biz_id=? and sys_id=?", bizid, sysid).Count(&count)
	return count
}
