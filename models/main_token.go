package models

//token相关
type MainToken struct {
	Symbol          string `gorm:"column:symbol;type:varchar(20);"`
	Name            string `gorm:"column:name;type:varchar(100);"`
	Decimals        int64  `gorm:"column:decimals;"`
	TotalSupply     string `gorm:"column:total_supply;type:varchar(100)"`
	Address         string `gorm:"column:address;type:varchar(100);primary_key"`
	CompiledVersion string `gorm:"column:compiled_version;type:varchar(50)"`
	BlockHeight     int64  `gorm:"column:block_height;type:bigint"`
	CreatorAddress  string `gorm:"column:creator_address;type:varchar(100)"`
	CreatorTime     int64  `gorm:"column:creator_time;"`
	Nonce           int32  `gorm:"column:nonce;"`
	Status          int32  `gorm:"column:status;"` //0=展示 1=不展示
	Model
}
