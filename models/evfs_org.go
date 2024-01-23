package models

// 组织
type EvfsOrg struct {
	OrgId           string `gorm:"column:org_id;type:varchar(100);primary_key"`
	OrgName         string `gorm:"column:org_name;type:varchar(45)"`
	TransactionHash string `gorm:"column:transaction_hash;type:varchar(100)"` //更改规则交易hash 默认null
	Model
}
