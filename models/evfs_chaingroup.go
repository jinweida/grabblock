package models

import (
	"github.com/jinzhu/gorm"
)

// 管理委员会组 ，默认初始化一些组
type EvfsChaingroup struct {
	GroupId         string `gorm:"column:group_id;type:varchar(100);primary_key"`
	Rule            int32  `gorm:"column:rule;"`         //当前规则
	PendingRule     int32  `gorm:"column:pending_rule;"` //待审核规则
	Name            string `gorm:"column:name;type:varchar(45)"`
	TransactionHash string `gorm:"column:transaction_hash;type:varchar(100)"` //更改规则交易hash 默认null
	Model
}

//规则申请

func (m *EvfsChaingroup) UpdatePendingRule(isEnd bool) error {
	if isEnd {
		m.Rule = m.PendingRule
	}
	return db.Save(&m).Error
}

//规则申请同意
func (m *EvfsChaingroup) UpdateRule(isEnd bool) error {
	return db.Model(&m).UpdateColumn("rule", gorm.Expr("pending_rule")).UpdateColumn("transaction_hash", m.TransactionHash).Error
}
