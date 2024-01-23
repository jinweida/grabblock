package models

//合约相关信息
type MainContract struct {
	ContractAddress  string `gorm:"column:contract_address;type:varchar(100);primary_key"`
	ContractName     string `gorm:"column:contract_name;type:varchar(50)"`
	CompiledVersion  string `gorm:"column:compiled_version;type:varchar(50)"`
	BlockHeight      int64  `gorm:"column:block_height;type:bigint"`
	CreatorAddress   string `gorm:"column:creator_address;type:varchar(100)"`
	CreatorTime      int64  `gorm:"column:creator_time;"`
	ContractCode     string `gorm:"column:contract_code;type:text;"`
	ContractAbi      string `gorm:"column:contract_abi;type:text;"`
	ContractBytecode string `gorm:"column:contract_bytecode;type:mediumtext;"`
	ContractExtdata  string `gorm:"column:contract_extdata;type:text;"` //发行是自定义
	Status           int32  `gorm:"column:status;"`                     //0=展示 1=不展示
	Model
}
