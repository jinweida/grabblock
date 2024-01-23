package service

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	constant "example.cn/grabblock/common"
	"example.cn/grabblock/conf"
	"example.cn/grabblock/entity"
	"example.cn/grabblock/models"
	"example.cn/grabblock/tools/log"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/golang/protobuf/proto"
)

type StratUp struct {
}

func NewStartup() *StratUp {
	return &StratUp{}
}

/*
*

	解析合约
*/
func (t *StratUp) ParsingContract() {
	start := time.Now()
	data := models.GetMainTransByContract(conf.Context.Node.Contractsize)
	for _, m := range data {
		trans := t.getTransactionInfo(m)
		//创世块解析
		if conf.ContractAddressSet.ContractAddress.ContractAdminAddress == m.StatusResult {
			log.Infof("解析创世块儿地址[%s]-ContractAdminAddress", m.StatusResult)
			admin := NewAdminGroup()
			admin_abi, err := abi.JSON(strings.NewReader(constant.AdminGroup))
			if err != nil {
				log.Errorf("AdminGroup is fail %s", err)
			}
			if status, err := admin.analysisGenesis_admingrop(trans, admin_abi); err == nil {
				models.UpdateIsdone(m.TransactionId, status)
			} else {
				log.Errorf("TransactionId=%s,error = %s", m.TransactionId, err.Error())
			}
		} else if conf.ContractAddressSet.ContractAddress.ContractCommitteeAddress == m.StatusResult {
			log.Infof("解析创世块儿地址[%s]-ContractCommitteeAddress", m.StatusResult)
			committeegroup := NewCommitteeGroup()
			committeegroup_abi, err := abi.JSON(strings.NewReader(constant.CommitteeGroup))
			if err != nil {
				log.Errorf("CommitteeGroup is fail %s", err)
			}
			if status, err := committeegroup.analysisGenesis_committee(trans, committeegroup_abi); err == nil {
				models.UpdateIsdone(m.TransactionId, status)
			} else {
				log.Errorf("TransactionId=%s,error = %s", m.TransactionId, err.Error())
			}
		} else {
			log.Infof("普通合约读取，交易hash[%s]", m.TransactionHash)
			if status, err := NewParsing().Decode(trans); err == nil {
				models.UpdateIsdone(m.TransactionId, status)
			} else {
				log.Errorf("TransactionId=%s,error = %s", m.TransactionId, err.Error())
			}
		}
	}
	log.Infof("parsing contract time cost=%v", time.Since(start))
}

/*
*

	解析上链数据（结构化、非结构化存储信息）
*/
func (t *StratUp) ParsingFile() {
	start := time.Now()
	data := models.GetMainTransByFile(conf.Context.Node.Contractsize)
	for _, m := range data {
		trans := t.getTransactionInfo(m)
		if status, err := t.DecodeFileOrData(trans); err == nil {
			models.UpdateIsdone(m.TransactionId, status)
		} else {
			log.Errorf("TransactionId=%s,error = %s", m.TransactionId, err.Error())
		}
	}
	log.Infof("parsing contract time cost=%v", time.Since(start))
}

/*
*

	解析上链数据
*/
func (t *StratUp) DecodeFileOrData(m *entity.TransactionInfo) (int64, error) {
	//获取codedata（proto）
	data, err := hex.DecodeString(m.Body.CodeData[2:])
	evfs := &entity.ContractEVFS{}
	err = proto.Unmarshal(data, evfs)
	if err != nil {
		log.Errorf("error decoding ContractEVFS: %v", err)
		return constant.PARSING_STATUS_FAIL, err
	}
	//文件上传处理(非机构化)
	if evfs.Function == entity.ContractEVFS_FILEUPLOADAPPLY {
		if evfs.FileUploadApplayData == nil {
			log.Errorf("evfs.FileUploadApplayData %v", err)
			return constant.PARSING_STATUS_FAIL, err
		}
		fileinfo := &entity.FileInfoData{}
		err = proto.Unmarshal(evfs.FileUploadApplayData.FileInfo, fileinfo)
		if err != nil {
			log.Errorf("error decoding evfs.FileUploadApplayData.FileInfo: %v", err)
			return constant.PARSING_STATUS_FAIL, err
		}
		evfsStorageFile := &models.EvfsStorageFile{
			FileSize:         fileinfo.FileSize,
			FileType:         fileinfo.FileType,
			SystemId:         "0x" + fileinfo.SysId,
			BizId:            "0x" + fileinfo.BizDomainId,
			CopyCount:        fileinfo.RepeatCount,
			SliceCount:       fileinfo.SliceCount,
			FileOwner:        string(evfs.FileUploadApplayData.FileOwner),
			SendFrontAddress: string(evfs.FileUploadApplayData.ClientAddr),
			TransactionHash:  m.Hash,
			ApplyTime:        m.Accepttimestamp,
			MarkDelete:       0, //0 否  1 是

		}
		//filehash 为空，代表是新加文件，不为空为修改文件
		if evfs.FileUploadApplayData.FileHash == nil {
			//将交易hash进行sha3运算，获得到文件hash唯一值
			filehashInsha3_256, _ := hex.DecodeString(m.Hash)
			evfsStorageFile.FileId = strings.ReplaceAll(crypto.Keccak256Hash(filehashInsha3_256).Hex(), "0x", "")
			evfsStorageFile.Op = 1 //添加
			evfsStorageFile.FileHash = evfsStorageFile.FileId
		} else {
			var build strings.Builder
			build.WriteString(string(evfs.FileUploadApplayData.FileHash))
			build.WriteString("_")
			nanosecond := time.Now().UnixNano() / 1e6
			build.WriteString(strconv.FormatInt(nanosecond, 10))
			evfsStorageFile.FileId = build.String()
			evfsStorageFile.Op = 2 //修改（修改文件，其实是新加了文件）
			evfsStorageFile.FileHash = string(evfs.FileUploadApplayData.FileHash)
		}
		if ext, ok := FileType[fileinfo.FileType]; ok {
			evfsStorageFile.ExtType = ext
		} else {
			evfsStorageFile.ExtType = constant.FILE_TYPE_OTHER
		}
		UploadTime := time.Unix(m.Status.Timestamp, 0)
		evfsStorageFile.Y = int32(UploadTime.Year())
		evfsStorageFile.D = int32(UploadTime.Day())
		evfsStorageFile.M = int32(UploadTime.Month())
		evfsStorageFile.H = int32(UploadTime.Hour())
		return constant.PARSING_DONE, evfsStorageFile.ApplyFile()
	} else if evfs.Function == entity.ContractEVFS_FILEUPLOADCONFIRM {
		//文件上传确认
		if evfs.ConfirmFileUplaodData == nil {
			log.Errorf("evfs.ConfirmFileUplaodData %v", err)
			return constant.PARSING_STATUS_FAIL, err
		}
		evfsStorageFile := &models.EvfsStorageFile{}
		//filehash := hexutil.Encode(evfs.ConfirmFileUplaodData.FileHash)
		fileversionhash := hexutil.Encode(evfs.ConfirmFileUplaodData.VersionHash)
		//更新文件状态
		return constant.PARSING_DONE, evfsStorageFile.ConfirmFile(fileversionhash, m.Accepttimestamp)
	} else if evfs.Function == entity.ContractEVFS_FILEDELETECONFIRM {
		//文件删除确认
		if evfs.FileDeleteConfirmData == nil {
			log.Errorf("evfs.FileDeleteConfirmData %v", err)
			return constant.PARSING_STATUS_FAIL, err
		}
		evfsStorageFile := &models.EvfsStorageFile{}
		//更新文件状态
		return constant.PARSING_DONE, evfsStorageFile.FileDeleteConfirmData(string(evfs.ConfirmFileUplaodData.VersionHash), m.Accepttimestamp)
	} else if evfs.Function == entity.ContractEVFS_DATASTORE {
		//结构化数据上传
		if evfs.DataStorage == nil {
			log.Errorf("evfs.DataStorage %v", err)
			return constant.PARSING_STATUS_FAIL, err
		}
		evfsStorageData := &models.EvfsStorageData{
			DataSize:        int64(len(evfs.DataStorage.Data)),
			SystemId:        m.Body.SysId,
			BizId:           m.Body.BizCode,
			DataOwner:       string(evfs.DataStorage.DataOwner),
			TransactionHash: m.Hash,
			ApplyTime:       m.Accepttimestamp,
			UplineTime:      m.Accepttimestamp,
			MarkDelete:      0, //0 否  1 是
			Approve:         1,
		}
		//RelDataHash 为空，代表是新加数据，不为空为修改数据
		if evfs.DataStorage.RelDataHash == nil {
			//将交易hash进行sha3运算，获得到文件hash唯一值
			filehashInsha3_256, _ := hex.DecodeString(m.Hash)
			evfsStorageData.DataId = strings.ReplaceAll(crypto.Keccak256Hash(filehashInsha3_256).Hex(), "0x", "")
			evfsStorageData.Op = 1 //添加
			evfsStorageData.DataHash = evfsStorageData.DataId
		} else {
			var build strings.Builder
			build.WriteString(string(evfs.DataStorage.RelDataHash))
			build.WriteString("_")
			nanosecond := time.Now().UnixNano() / 1e6
			build.WriteString(strconv.FormatInt(nanosecond, 10))
			evfsStorageData.DataId = build.String()
			evfsStorageData.Op = 2 //修改（修改文件，其实是新加了文件）
			evfsStorageData.DataHash = string(evfs.DataStorage.RelDataHash)
		}
		UploadTime := time.Unix(m.Status.Timestamp, 0)
		evfsStorageData.Y = int32(UploadTime.Year())
		evfsStorageData.D = int32(UploadTime.Day())
		evfsStorageData.M = int32(UploadTime.Month())
		evfsStorageData.H = int32(UploadTime.Hour())
		return constant.PARSING_DONE, evfsStorageData.ApplyData()
	} else if evfs.Function == entity.ContractEVFS_DATASTOREDEL {
		//结构化数据删除
		if evfs.DataDelData == nil {
			log.Errorf("evfs.DataDelData %v", err)
			return constant.PARSING_STATUS_FAIL, err
		}
		evfsStorageData := &models.EvfsStorageData{}
		datahash := hexutil.Encode(evfs.DataDelData.DataHash)
		//更新文件状态(标记删除)
		return constant.PARSING_DONE, evfsStorageData.DataDeleteConfirm(strings.ReplaceAll(datahash, "0x", ""), m.Accepttimestamp)
	} else {
		log.Errorf("evfs.FileUploadApplayData %v", errors.New(fmt.Sprintf("evfs.Function状态=%s", evfs.Function)))
		return constant.PARSING_NOT, nil
	}
}
func (t *StratUp) getTransactionInfo(m models.MainTransaction) *entity.TransactionInfo {
	trans := &entity.TransactionInfo{}
	trans.Hash = m.TransactionHash
	trans.Status = &entity.TransactionStatus{}
	trans.Status.Status = m.Status
	trans.Status.Height = m.BlockHeight
	trans.Status.Hash = m.BlockHash
	trans.Status.Result = m.StatusResult
	trans.Status.Timestamp = m.StatusTime
	trans.Accepttimestamp = m.AcceptTime
	trans.Body = &entity.TransactionBody{}
	trans.Body.Address = m.FromAddress
	trans.Body.Nonce = m.Nonce
	trans.Body.CodeData = m.CodeData
	trans.Body.InnerCodetype = m.InnerCodetype
	trans.Body.Timestamp = m.Timestamp
	trans.Body.Outputs = make([]*entity.TransactionOutput, 0)
	trans.Body.SysId = m.SysId
	// trans.Body.BizId = m.BizId
	output := &entity.TransactionOutput{
		Address: m.ToAddress,
		Amount:  m.Amount,
	}
	trans.Body.Outputs = append(trans.Body.Outputs, output)
	return trans
}
