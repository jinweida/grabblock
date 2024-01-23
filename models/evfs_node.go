package models

import (
	"fclink.cn/grabblock/tools/log"
)

/***
chainNodeApply(address,bytes,address,bytes,uint256,uint256,uint256)
企业地址<br>企业名称<br>节点地址<br>节点信息<br>初始Token数量<br>节点类型: 1:记账节点 2:只读节点 3:前置节点<br>操作类型: 1:添加 2:移除
*/
// 记账节点/同步节点/客户端节点
type EvfsNode struct {
	ChainnodeId     string `gorm:"column:chainnode_id;type:varchar(100);primary_key"` //代办ID
	TransactionHash string `gorm:"column:transaction_hash;type:varchar(100);"`       //交易hash
	OrgId           string `gorm:"column:org_id;type:varchar(100);"`                  //企业address
	OrgName         string `gorm:"column:org_name;type:varchar(100);"`                //企业名称
	NodeAddress     string `gorm:"column:node_address;type:varchar(100);"`            //节点address
	NodeInfo        string `gorm:"column:node_info;type:varchar(500);"`               //节点信息 扩展
	Cpu             string `gorm:"column:cpu;type:varchar(100);"`                     //由nodeinfo 解析而来
	Memory          string `gorm:"column:memory;type:varchar(100);"`                  //由nodeinfo 解析而来
	Disk            string `gorm:"column:disk;type:varchar(100);"`                    //由nodeinfo 解析而来
	Bandwidth       string `gorm:"column:bandwidth;type:varchar(100);"`               //由nodeinfo 解析而来
	CapaditySize    int64  `gorm:"column:capadity_size;"`                             //结构化数据容量
	FileCapaditySize int64  `gorm:"column:file_capadity_size;"`                    //非结构化数据容量
	NodeType        int32  `gorm:"column:node_type;"`                                 //节点类型 1:记账节点 2:只读节点 3:前置节点
	Approve         int32  `gorm:"column:approve;"`                                   //审批状态 0=审核中 1=同意 2=拒绝
	Op              int64  `gorm:"column:op;"`                                        //操作 1 添加 2移除
	Model
}

func (m *EvfsNode) Add() error {
	tx := db.Begin()
	//写入org
	if _, ok := org[m.OrgId]; !ok {
		org[m.OrgId] = &EvfsOrg{
			OrgId:           m.OrgId,
			OrgName:         m.OrgName,
			TransactionHash: m.TransactionHash,
		}

		if err := tx.Create(org[m.OrgId]).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	if err := tx.Create(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

////更新数据
//func (m *EvfsNode) UpdateApprove() error {
//	data := make([]EvfsNode, 0)
//	db.Where("req_id=?",m.ReqId).Find(&data)
//	if len(data)>0 && data[0].Op==2 {
//		//删除同意
//		return db.Model(&m).Where("node_address=?", data[0].NodeAddress).Delete(m).Error
//	}else{
//		//添加同意
//		return db.Model(&m).Where("req_id=?", m.ReqId).UpdateColumn("approve", m.Approve).Error
//	}
//}
//
////删除数据
//func (m *EvfsNode) Delete() error {
//	tx := db.Begin()
//	log.Infof("------------节点删除：address [%s]",m.NodeAddress)
//	if err := tx.Where("node_address=?",m.NodeAddress).Delete(&m).Error; err != nil {
//		tx.Rollback()
//		return err
//	}
//	return tx.Commit().Error
//}


func (m *EvfsNode) Delete() error {
	tx := db.Begin()
	log.Infof("------------节点删除：address [%s]",m.NodeAddress)
	if err := tx.Where("node_address=?",m.NodeAddress).Delete(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (m *EvfsNode) UpdateApprove() error {
	data := make([]EvfsResourcenode, 0)
	db.Where("chainnode_id=?",m.ChainnodeId).Find(&data)
	if len(data)>0 && data[0].Op==2 {
		return db.Model(&m).Where("node_address=?", data[0].NodeAddress).Delete(m).Error
	}else{
		return db.Model(&m).Where("chainnode_id=?", m.ChainnodeId).UpdateColumn("approve", m.Approve).UpdateColumn("transaction_hash", m.TransactionHash).Error
	}
}

/**
  分页查询 主节点数据
 */
func (m *EvfsNode) Find(Page int,PageSize int) []EvfsNode {
	data := make([]EvfsNode, 0)
	if Page > 0 && PageSize > 0 {
		db.Order("created asc").Where("approve=1 and node_type=1").Limit(PageSize).Offset((Page - 1) * PageSize).Find(&data)
	}else{
		db.Order("created asc").Where("approve=1 and node_type=1").Limit(10).Offset(0).Find(&data)
	}
	return data
}

/**
   计算主节点总数
 */
func (m *EvfsNode) Count() int {
	count := 0
	db.Model(EvfsNode{}).Where("approve=1 and node_type=1").Count(&count)
	return count
}
