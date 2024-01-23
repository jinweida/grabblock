package models

import "fclink.cn/grabblock/tools/log"

//存储节点
type EvfsResourcenode struct {
	ReqId          string `gorm:"column:req_id;type:varchar(100);primary_key"` //待办ID
	NodeAddress     string `gorm:"column:node_address;type:varchar(100);"`      //节点address
	NodeName     string `gorm:"column:node_name;type:varchar(100);"`      //节点名称
	TransactionHash string `gorm:"column:transaction_hash;type:varchar(100);"`  //交易hash
	OrgId           string `gorm:"column:org_id;type:varchar(100);"`            //组织ID
	OrgName        string `gorm:"column:org_name;type:varchar(100);"`          //组织名称
	DomainId        string `gorm:"column:domain_id;type:varchar(100);"`         //存储域ID
	NodeInfo        string `gorm:"column:node_info;type:varchar(500);"`         //节点信息 扩展
	Url             string `gorm:"column:url;type:varchar(500);"`               //节点URL
	Cpu             string `gorm:"column:cpu;type:varchar(100);"`               //由nodeinfo 解析而来
	Memory          string `gorm:"column:memory;type:varchar(100);"`            //由nodeinfo 解析而来
	Disk            string `gorm:"column:disk;type:varchar(100);"`              //由nodeinfo 解析而来
	Bandwidth       string `gorm:"column:bandwidth;type:varchar(100);"`         //由nodeinfo 解析而来
	CapaditySize    int64  `gorm:"column:capadity_size;"`                       //结构化数据总容量
	FileCapaditySize int64  `gorm:"column:file_capadity_size;"`                 //非结构化数据容量
	Approve         int32  `gorm:"column:approve;"`                             //审批状态 0=审核中 1=同意 2=拒绝
	Op              int64  `gorm:"column:op;"`                                   //操作 1 添加 2移除
	Model
}

//存储节点申请
func (m *EvfsResourcenode) Add() error {
	return db.Create(&m).Error
}

//更新数据
func (m *EvfsResourcenode) UpdateApprove() error {
	data := make([]EvfsResourcenode, 0)
	db.Where("req_id=?",m.ReqId).Find(&data)
	if len(data)>0 && data[0].Op==2 {
		//删除同意
		return db.Model(&m).Where("node_address=?", data[0].NodeAddress).Delete(m).Error
	}else{
		//添加同意
		return db.Model(&m).Where("req_id=?", m.ReqId).UpdateColumn("approve", m.Approve).Error
	}
}

//删除数据
func (m *EvfsResourcenode) Delete() error {
	tx := db.Begin()
	log.Infof("------------节点删除：address [%s]",m.NodeAddress)
	if err := tx.Where("node_address=?",m.NodeAddress).Delete(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}


/**
  分页查询 资源节点数据
*/
func (m *EvfsResourcenode) Find(Page int,PageSize int) []EvfsResourcenode {
	data := make([]EvfsResourcenode, 0)
	if Page > 0 && PageSize > 0 {
		db.Order("created asc").Where("approve=1").Limit(PageSize).Offset((Page - 1) * PageSize).Find(&data)
	}else{
		db.Order("created asc").Where("approve=1").Limit(10).Offset(0).Find(&data)
	}
	return data
}

/**
  计算资源节点总数
*/
func (m *EvfsResourcenode) Count() int {
	count := 0
	db.Model(EvfsResourcenode{}).Where("approve=1").Count(&count)
	return count
}