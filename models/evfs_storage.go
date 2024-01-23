package models

import (
	"time"

	constant "example.cn/grabblock/common"
	"example.cn/grabblock/tools/log"
	"github.com/jinzhu/gorm"
)

// 数据存管域
type EvfsStorage struct {
	ReqId               string              `gorm:"column:req_id;type:varchar(100);primary_key"`    //代办ID
	DomainId            string              `gorm:"column:domain_id;type:varchar(100);"`            //存储域ID
	StorageName         string              `gorm:"column:storage_name;type:varchar(200);"`         //存储域名称
	CapacityReqId       string              `gorm:"column:capadity_req_id;type:varchar(100);"`      //license 代办ID
	CapacitySize        int64               `gorm:"column:capadity_size;"`                          //初始容量
	PendingCapacitySize int64               `gorm:"column:pending_capadity_size;"`                  //初始容量
	UsedFileSize        int64               `gorm:"column:used_file_size;"`                         //企业address
	UsedDataSize        int64               `gorm:"column:used_data_size;"`                         //企业address
	FileCount           int64               `gorm:"column:file_count;"`                             //企业address
	DataCount           int64               `gorm:"column:data_count;"`                             //企业address
	OrgNum              int32               `gorm:"column:org_num;"`                                //企业address
	SysNum              int32               `gorm:"column:sys_num;"`                                //企业address
	UserNum             int32               `gorm:"column:user_num;"`                               //企业address
	OutChainNum         int32               `gorm:"column:out_chain_num;"`                          //企业address
	NodeNum             int32               `gorm:"column:node_num;"`                               //企业address
	FileSize            int64               `gorm:"column:file_size;"`                              //企业address
	DataSize            int64               `gorm:"column:data_size;"`                              //企业address
	ClientNum           int32               `gorm:"column:client_num;"`                             //企业address
	CommitteeNum        int32               `gorm:"column:committee_num;"`                          //企业address
	RuleReqId           string              `gorm:"column:rule_req_id;type:varchar(100);"`          //规则代办ID
	Rule                int32               `gorm:"column:rule;"`                                   //当前规则
	PendingRule         int32               `gorm:"column:pending_rule;"`                           //待审核规则
	OrgId               string              `gorm:"column:org_id;type:varchar(100);"`               //企业address
	OrgName             string              `gorm:"column:org_name;type:varchar(100);"`             //企业名称
	Approve             int32               `gorm:"column:approve;"`                                //审批状态 0=审核中 1=同意 2=拒绝
	TransactionHash     string              `gorm:"column:transaction_hash;type:varchar(100)"`      //更改规则交易hash 默认null
	RuleTransactionHash string              `gorm:"column:rule_transaction_hash;type:varchar(100)"` //更改规则交易hash 默认null
	SizeTransactionHash string              `gorm:"column:size_transaction_hash;type:varchar(100)"` //更改规则交易hash 默认null
	Members             []EvfsStorageMember `gorm:"-"`
	AcceptTime          int64               `gorm:"column:accept_time;"`
	Model
}

func (m *EvfsStorage) Add() error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		log.Errorf("err:%s", err.Error())
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
			tx.Rollback()
			log.Errorf("err:%s", err.Error())
			return err
		}
	}

	if err := tx.Create(&m).Error; err != nil {
		tx.Rollback()
		log.Errorf("err:%s", err.Error())
		return err
	}

	for _, member := range m.Members {

		if err := tx.Create(&member).Error; err != nil {
			tx.Rollback()
			log.Errorf("err:%s", err.Error())
			return err
		}
	}
	evfsStorageLicens := &EvfsStorageLicens{
		TransactionHash: m.TransactionHash,
		CapacitySize:    m.CapacitySize,
		AcceptTime:      m.AcceptTime,
		ReqId:           m.ReqId,
		Approve:         constant.APPROVE_ING,
	}
	accept := time.Unix(m.AcceptTime, 0)
	evfsStorageLicens.Y = int32(accept.Year())
	evfsStorageLicens.M = int32(accept.Month())
	evfsStorageLicens.D = int32(accept.Day())
	evfsStorageLicens.H = int32(accept.Hour())

	if err := tx.Create(&evfsStorageLicens).Error; err != nil {
		tx.Rollback()
		log.Errorf("err:%s", err.Error())
		return err
	}
	return tx.Commit().Error
}

// 规则申请同意
func (m *EvfsStorage) UpdateRule(isEnd bool) error {
	return db.Model(&m).Where("domain_id", m.DomainId).Where("rule_req_id", m.RuleReqId).UpdateColumn("rule", gorm.Expr("pending_rule")).UpdateColumn("rule_transaction_hash", m.RuleTransactionHash).Error
}

// 规则申请同意
func (m *EvfsStorage) UpdatePendingRule(isEnd bool) error {
	sql := db.Model(&m).Where("domain_id", m.DomainId)
	if isEnd {
		sql = sql.UpdateColumn("rule", m.PendingRule)
	} else {
		sql = sql.UpdateColumn("pending_rule", m.PendingRule)
	}
	sql = sql.UpdateColumn("rule_req_id", m.RuleReqId)
	return sql.UpdateColumn("rule_transaction_hash", m.RuleTransactionHash).Error
}

// 存储License申请
func (m *EvfsStorage) UpdatePendingSize(isEnd bool) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		log.Errorf("err:%s", err.Error())
		return err
	}
	sql := tx.Model(&m).Where("domain_id=?", m.DomainId)
	if isEnd {
		sql = sql.UpdateColumn("capadity_size", gorm.Expr("capadity_size + ?", m.PendingCapacitySize)).UpdateColumn("size_transaction_hash", m.SizeTransactionHash)
	} else {
		sql = sql.UpdateColumn("pending_capadity_size", m.PendingCapacitySize).UpdateColumn("size_transaction_hash", m.SizeTransactionHash)
	}
	sql = sql.UpdateColumn("capadity_req_id", m.CapacityReqId)
	if err := sql.Error; err != nil {
		tx.Rollback()
		log.Errorf("err:%s", err.Error())
		return err
	}
	evfsStorageLicens := &EvfsStorageLicens{
		TransactionHash: m.SizeTransactionHash,
		CapacitySize:    m.PendingCapacitySize,
		AcceptTime:      m.AcceptTime,
		ReqId:           m.ReqId,
		Approve:         constant.APPROVE_ING,
	}
	accept := time.Unix(m.AcceptTime, 0)
	evfsStorageLicens.Y = int32(accept.Year())
	evfsStorageLicens.M = int32(accept.Month())
	evfsStorageLicens.D = int32(accept.Day())
	evfsStorageLicens.H = int32(accept.Hour())
	if isEnd {
		evfsStorageLicens.Approve = constant.APPROVE_SUCCESS
	}
	if err := tx.Create(&evfsStorageLicens).Error; err != nil {
		tx.Rollback()
		log.Errorf("err:%s", err.Error())
		return err
	}
	return tx.Commit().Error
}

// 存储License通过
func (m *EvfsStorage) UpdateSize() error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		log.Errorf("err:%s", err.Error())
		return err
	}
	sql := tx.Model(&m).Where("capadity_req_id=?", m.CapacityReqId)
	sql = sql.UpdateColumn("capadity_size", gorm.Expr("capadity_size + pending_capadity_size")).UpdateColumn("size_transaction_hash", m.SizeTransactionHash)
	if err := sql.Error; err != nil {
		tx.Rollback()
		log.Errorf("err:%s", err.Error())
		return err
	}
	//licents_logs
	if err := tx.Model(&EvfsStorageLicens{}).Where("req_id=?", m.CapacityReqId).UpdateColumn("approve", m.Approve).UpdateColumn("transaction_hash", m.SizeTransactionHash).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

// 审核
func (m *EvfsStorage) UpdateApprove() error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		log.Errorf("err:", err.Error())
		return err
	}
	if err := tx.Model(&m).UpdateColumn("approve", m.Approve).UpdateColumn("domain_id", m.DomainId).UpdateColumn("transaction_hash", m.TransactionHash).Error; err != nil {
		tx.Rollback()
		return err
	}
	//storage_member
	if err := tx.Model(&EvfsStorageMember{}).Where("req_id=?", m.ReqId).UpdateColumn("approve", m.Approve).UpdateColumn("transaction_hash", m.TransactionHash).UpdateColumn("domain_id", m.DomainId).Error; err != nil {
		tx.Rollback()
		return err
	}
	//licents_logs
	if err := tx.Model(&EvfsStorageLicens{}).Where("req_id=?", m.ReqId).UpdateColumn("approve", m.Approve).UpdateColumn("transaction_hash", m.TransactionHash).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
