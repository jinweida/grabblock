package models

import (
	"errors"
	"fmt"
)

/***
  业务域下----解析合约信息（相关abi=createContractApply、createContractAgree、contractEnableApply、contractEnableAgree）
*/
type EvfsBizContractFlow struct {
	ReqId                  string `gorm:"column:req_id;type:varchar(100);primary_key"`                 //待办ID
	ContractAddress        string `gorm:"column:contract_address;type:varchar(100)"`                   //合约地址
	BizId                  string `gorm:"column:biz_id;type:varchar(100);"`                            //业务域ID
	ContractName           string `gorm:"column:contract_name;type:varchar(100);"`                     //合约名称
	TransactionHash        string `gorm:"column:transaction_hash;type:varchar(100);"`                  //交易hash
	ContractMark           string `gorm:"column:contract_mark;type:varchar(100);"`                     //合约备注
	StatusInfo             int64    `gorm:"column:status_info;"`                                       //是否冻结  0 创建 1 冻结 2 解冻
	FrozenTime             int    `gorm:"column:frozen_time;"`                                         //冻结时间
	Approve                int    `gorm:"column:approve;"`                                             //审批状态 0=审核中 1=同意 2=拒绝
	ContractInfo           string `gorm:"column:contract_info;type:varchar(100);"`                     //合约信息
	Model
}

//合约创建
func (m *EvfsBizContractFlow) CreateContractAdd(releasetime int64) error {
	//同意--直接记录合约信息
	if m.Approve == 1{
		addEvfsBizContractRecord(m,m.ContractName,m.ContractMark,releasetime)
	}
	//记录合约流水
	return db.Create(&m).Error
}

//合约同意创建
func (m *EvfsBizContractFlow) CreateContractUpdateApprove(releasetime int64) error {
	//查找出同意添加的业务合约，然后入库至EvfsBizContractRecord
	data := make([]EvfsBizContractFlow, 0)
	db.Where("req_id=?",m.ReqId).Find(&data)
	if len(data) < 0 {
		return errors.New(fmt.Sprintf("hash=%s,CreateContractUpdateApprove 未找到需更新数据，可能抓块有数据丢失。可以选择重新抓块。", m.TransactionHash))
	}
	//记录新创建的合约信息
	addEvfsBizContractRecord(m,data[0].ContractName,data[0].ContractMark,releasetime)
	//添加同意-更新合约流水
	return db.Model(&m).Where("req_id=?", m.ReqId).UpdateColumn("approve", m.Approve).UpdateColumn("contract_address", m.ContractAddress).Error
}

//添加合约记录（用于后续页面API查询）
func addEvfsBizContractRecord(m *EvfsBizContractFlow,contractName string,contractMark string,releasetime int64)  {
	evfsBizContractRecord := EvfsBizContractRecord{
		ContractAddress:      m.ContractAddress,         //合约地址
		BizId:                m.BizId,                   //业务域ID
		ContractName:         contractName,              //合约名称
		ContractMark:         contractMark,              //合约备注
		IsFrozen:             int(1),                    //1 正常
		ReleaseTime:          releasetime,               //发布时间（时间戳）
		InureTime:            releasetime,               //生效时间（时间戳）
	}
	evfsBizContractRecord.Add()
}

//合约冻结、解冻
func (m *EvfsBizContractFlow) ContractFrozen(inuretime int64) error{
   if m.Approve == 1 {
   	   evfsBizContractRecord := EvfsBizContractRecord{}
   	   // 冻结 或 解冻合约--更新record
   	     // 1 冻结
   	  if m.StatusInfo==1{
		  return db.Model(&evfsBizContractRecord).Where("contract_address=? and biz_id=?", m.ContractAddress,m.BizId).UpdateColumn("is_frozen", 2).UpdateColumn("inuretime",inuretime).Error
	  }else{
	  	//解冻
		  return db.Model(&evfsBizContractRecord).Where("contract_address=? and biz_id=?", m.ContractAddress,m.BizId).UpdateColumn("is_frozen", 1).UpdateColumn("inuretime",inuretime).Error
	  }
   }
	//记录合约流水
	return db.Create(&m).Error
}

//合约冻结、解冻-同意操作
func (m *EvfsBizContractFlow) ContractFrozenUpdateApprove(inuretime int64) error {
	//查询待办ID信息，判断是冻结操作或解冻操作
	data := make([]EvfsBizContractFlow, 0)
	db.Where("req_id=?",m.ReqId).Find(&data)
	if len(data) < 0 {
		return errors.New(fmt.Sprintf("hash=%s,ContractFrozenUpdateApprove 未找到需更新数据，可能抓块有数据丢失。可以选择重新抓块。", m.TransactionHash))
	}
	evfsBizContractRecord := EvfsBizContractRecord{}
	//更新record合约信息
	//1 冻结
	if data[0].StatusInfo==1{
		return db.Model(&evfsBizContractRecord).Where("contract_address=? and biz_id=?", m.ContractAddress,m.BizId).UpdateColumn("is_frozen", 2).UpdateColumn("inuretime",inuretime).Error
	}else{
		//解冻
		return db.Model(&evfsBizContractRecord).Where("contract_address=? and biz_id=?", m.ContractAddress,m.BizId).UpdateColumn("is_frozen", 1).UpdateColumn("inuretime",inuretime).Error
	}
	//添加同意-更新合约流水
	return db.Model(&m).Where("req_id=?", m.ReqId).UpdateColumn("approve", m.Approve).UpdateColumn("transaction_hash", m.TransactionHash).Error
}