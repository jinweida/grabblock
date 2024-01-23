package models

import (
	"example.cn/grabblock/tools/log"
)

type MainAccountBalance struct {
	Id           int64  `gorm:"column:id;type:bigint;AUTO_INCREMENT;primary_key"`
	Address      string `gorm:"column:address;type:varchar(100)"`
	TokenSymbol  string `gorm:"column:token_symbol;type:varchar(20)"` //币种
	TokenAddress string `gorm:"column:token_address;type:varchar(100)"`
	Balance      string `gorm:"column:balance;type:varchar(100)"`
	CreatedBy    string `gorm:"column:created_by;type:varchar(32)"`
	ModifiedBy   string `gorm:"column:modified_by;type:varchar(32)"`
	Remark       string `gorm:"column:remark;type:varchar(255)"` //备注
	CasVersion   int32  `gorm:"column:cas_version;"`             //乐观锁
	Tag          string `gorm:"column:tag;type:varchar(32)"`     //标签
	AccountType  int32  `gorm:"column:account_type;"`            //账户类型:1-主链资产交易,2-720token
}

func AddMainAccountBalance(list []*MainAccountBalance) error {
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
	for _, tran := range list {
		var count int32
		tx.Model(&tran).Where("address=? and token_address=?", tran.Address, tran.TokenAddress).Count(&count)
		if count == 0 {
			if err := tx.Create(tran).Error; err != nil {
				tx.Rollback()
				log.Errorf("insert balance err:%s", err.Error())
				return err
			}
		} else {
			if err := tx.Model(&tran).Where("address=? and token_address=?", tran.Address, tran.TokenAddress).UpdateColumn("balance", tran.Balance).Error; err != nil {
				tx.Rollback()
				log.Errorf("update balance err:%s", err.Error())
				return err
			}
		}
	}
	return tx.Commit().Error
}
