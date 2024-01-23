package models

import (
	"fclink.cn/grabblock/tools/log"
	"github.com/jinzhu/gorm"
)

type EvfsBiz struct {
	ReqId               string          `gorm:"column:req_id;type:varchar(100);primary_key"`    //代办ID
	DomainId            string          `gorm:"column:domain_id;type:varchar(100);"`            //存储域ID
	BizId               string          `gorm:"column:biz_id;type:varchar(100);"`               //业务域ID
	BizName             string          `gorm:"column:biz_name;type:varchar(200);"`             //业务域信息
	OrgId               string          `gorm:"column:org_id;type:varchar(100);"`               //企业address
	OrgName             string          `gorm:"column:org_name;type:varchar(100);"`             //企业名称
	Approve             int32           `gorm:"column:approve;"`                                //审批状态 0=审核中 1=同意 2=拒绝
	TransactionHash     string          `gorm:"column:transaction_hash;type:varchar(100)"`      //更改规则交易hash 默认null
	RuleReqId           string          `gorm:"column:rule_req_id;type:varchar(100);"`          //规则代办ID
	Rule                int32           `gorm:"column:rule;"`                                   //当前规则
	PendingRule         int32           `gorm:"column:pending_rule;"`                           //待审核规则
	RuleTransactionHash string          `gorm:"column:rule_transaction_hash;type:varchar(100)"` //更改规则交易hash 默认null
	Members             []EvfsBizMember `gorm:"-"`
	FileCount           int32           `gorm:"-"`
	DataCount           int32           `gorm:"-"`
	Model
}

func (m *EvfsBiz) Add(agree bool) error {
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
			tx.Rollback()
			return err
		}
	}

	if err := tx.Create(&m).Error; err != nil {
		tx.Rollback()
		return err
	}

	for _, member := range m.Members {
		if err := tx.Create(&member).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	//如果同意-入库统计信息（用于后续数据统计）
	if agree {
		evfsBizStatic := EvfsBizStatic{
			BizId: m.BizId,
			BizName: m.BizName,
		}
		evfsBizStatic.Add()
	}
	return tx.Commit().Error
}

func (m *EvfsBiz) UpdateApprove() error {
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
	if err := tx.Model(&m).Where("req_id=?", m.ReqId).UpdateColumn("approve", m.Approve).UpdateColumn("biz_id", m.BizId).UpdateColumn("transaction_hash", m.TransactionHash).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(&EvfsBizMember{}).Where("req_id=?", m.ReqId).UpdateColumn("approve", m.Approve).UpdateColumn("transaction_hash", m.TransactionHash).UpdateColumn("biz_id", m.BizId).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

//规则申请同意
func (m *EvfsBiz) UpdateRule(isEnd bool) error {
	return db.Model(&m).Where("biz_id", m.DomainId).Where("rule_req_id", m.RuleReqId).UpdateColumn("rule", gorm.Expr("pending_rule")).UpdateColumn("rule_transaction_hash", m.RuleTransactionHash).Error
}

//规则申请同意
func (m *EvfsBiz) UpdatePendingRule(isEnd bool) error {
	db = db.Model(&m).Where("biz_id", m.BizId)
	if isEnd {
		db = db.UpdateColumn("rule", m.PendingRule)
	} else {
		db = db.UpdateColumn("pending_rule", m.PendingRule)
	}
	db = db.UpdateColumn("rule_req_id", m.RuleReqId)
	return db.UpdateColumn("rule_transaction_hash=?", m.RuleTransactionHash).Error
}


func (m *EvfsBiz) UseBizidFind(bizid string) EvfsBiz {
	data := make([]EvfsBiz, 0)
	db.Where("biz_id=?",bizid).Find(&data)
	return data[0]
}
