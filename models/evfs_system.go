package models

import (
	"fmt"
	"github.com/pkg/errors"
	"strings"
)

type EvfsSystem struct {
	ReqId           string `gorm:"column:req_id;type:varchar(100);primary_key"`
	OrgId           string `gorm:"column:org_id;type:varchar(100);"` //企业address
	OrgName         string `gorm:"column:org_name;type:varchar(45)"`
	BizId           string `gorm:"column:biz_id;type:varchar(100);"`          //业务域ID
	SysId           string `gorm:"column:sys_id;type:varchar(100);"`          //系统ID
	SysName         string `gorm:"column:sys_name;type:varchar(100);"`        //系统名称
	SysInfo         string `gorm:"column:sys_info;type:varchar(100);"`        //系统信息
	Approve         int32  `gorm:"column:approve;"`                           //审批状态 0=审核中 1=同意 2=拒绝
	TransactionHash string `gorm:"column:transaction_hash;type:varchar(100)"` //更改规则交易hash 默认null
	Model
}

func (m *EvfsSystem) Add(agree bool) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	//写入org
	if _, ok := org[m.OrgId]; !ok {
		org[m.OrgId] = &EvfsOrg{
			OrgId:           m.OrgId,
			OrgName:         m.OrgName,
			TransactionHash: m.TransactionHash,
		}
		if err := tx.Create(org[m.OrgId]).Error; err != nil {
			if  strings.Contains(err.Error(),"Duplicate"){
				//有可能企业重复，重复认为是更新
				return db.Model(org[m.OrgId]).Where("org_id=?", m.OrgId).UpdateColumn("org_name", m.OrgName).UpdateColumn("transaction_hash", m.TransactionHash).Error
			}else {
				tx.Rollback()
			}
		}
	}
    //新增业务系统时，增加业务系统统计数据
	if agree {
		evfsSystemStatic := EvfsSystemStatic{
			BizId: m.BizId,  //业务域ID
			OrgId: m.OrgId,  //企业ID
			OrgName:m.OrgName, //企业名称
			SysId:m.SysId,    //业务系统ID
			SysName:m.SysName, //业务系统名称
		}
		evfsSystemStatic.Add()
	}

	if err := tx.Create(&m).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (m *EvfsSystem) UpdateApprove() error {
	//通过待办ID查找业务系统信息
	sysinfo,count := m.UseReqidFind(m.ReqId)
	if count==1{
		evfsSystemStatic := EvfsSystemStatic{
			BizId: sysinfo.BizId,  //业务域ID
			OrgId: sysinfo.OrgId,  //企业ID
			OrgName:sysinfo.OrgName, //企业名称
			SysId:sysinfo.SysId,    //业务系统ID
			SysName:sysinfo.SysName, //业务系统名称
		}
		//增加统计信息，后续会根据统计表信息，统计此表下相关数据
		evfsSystemStatic.Add()
		return db.Model(&m).Where("req_id=?", m.ReqId).UpdateColumn("approve", m.Approve).UpdateColumn("transaction_hash", m.TransactionHash).Error
	}else{
		return errors.New(fmt.Sprintf("reqid=%s", m.ReqId))
	}
}

func (m *EvfsSystem) CountOrg(bizid string) int64 {
	count := int64(0)
	db.Model(m).Select("DISTINCT(org_id)").Where("biz_id=?",bizid).Count(&count)
	return count
}

func (m *EvfsSystem) CountSys(bizid string) int64 {
	count := int64(0)
	db.Model(m).Select("DISTINCT(sys_id)").Where("biz_id=?",bizid).Count(&count)
	return count
}


func (m *EvfsSystem) UseReqidFind(req_id string) (EvfsSystem,int64) {
	data := make([]EvfsSystem, 0)
	db.Where("req_id=?",req_id).Find(&data)
	if len(data)==0 {
		kk := EvfsSystem{}
		return kk,0
	}else{
		return data[0],1
	}
}

