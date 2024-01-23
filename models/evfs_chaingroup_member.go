package models

import (
	"encoding/json"
	"os"
)

//管理委员会成员
type EvfsChaingroupMember struct {
	ChaincommitteeId string    `gorm:"column:chaincommittee_id;type:varchar(100);primary_key"` //代办ID
	MainChainGroupId string    `gorm:"column:main_chain_group_id;type:varchar(45)"`            //1 管理员  2委员会
	MemberAddress    string    `gorm:"column:member_address;type:varchar(100)"`                //成员address
	JoinTime         int64     `gorm:"column:jsoin_time;"`                                     //加入日期
	Approve          int32     `gorm:"column:approve;"`                                        //审批状态 0=审核中 1=同意 2=拒绝
	TransactionHash  string    `gorm:"column:transaction_hash;type:varchar(100)"`              //更改规则交易hash 默认null
	MemberName       string    `gorm:"column:member_name;type:varchar(100)"`                   //成员名称
	Model
}

//成员申请
func (m *EvfsChaingroupMember) Add() error {
	return db.Create(&m).Error
}

//成员同意
func (m *EvfsChaingroupMember) Delete() error {
	return db.Delete(&m).Error
}

//成员同意
func (m *EvfsChaingroupMember) UpdateApprove() error {
	json.NewEncoder(os.Stdout).Encode(m)
	err := db.Model(&m).Update("transaction_hash", m.TransactionHash).Update("approve", m.Approve).Error

	return err
}
