package models

/***
  业务域下----合约信息
*/
type EvfsBizContractRecord struct {
	ContractAddress        string `gorm:"column:contract_address;type:varchar(100);primary_key"`       //合约地址
	BizId                  string `gorm:"column:biz_id;type:varchar(100);"`                            //业务域ID
	ContractName           string `gorm:"column:contract_name;type:varchar(100);"`                     //合约名称
	ContractMark           string `gorm:"column:contract_mark;type:varchar(100);"`                     //合约备注
	IsFrozen               int    `gorm:"column:is_frozen;"`                                           //是否冻结  1 正常 2 冻结
	InureTime              int64  `gorm:"column:inure_time;"`                                         //生效时间
	CallCount              int    `gorm:"column:call_count;"`                                          //调用次数
	ReleaseTime            int64  `gorm:"column:release_time;"`                                        //发布时间
	Model
}

//添加合约
func (m *EvfsBizContractRecord) Add() error {
	return db.Create(&m).Error
}

/**
  分页查询 某业务域下数据
*/
func (m *EvfsBizContractRecord) Find(bizid string,Page int,PageSize int) []EvfsBizContractRecord {
	data := make([]EvfsBizContractRecord, 0)
	if Page > 0 && PageSize > 0 {
		db.Order("created asc").Where("biz_id=?",bizid).Limit(PageSize).Offset((Page - 1) * PageSize).Find(&data)
	}else{
		db.Order("created asc").Where("biz_id=?",bizid).Limit(10).Offset(0).Find(&data)
	}
	return data
}

/**
  某业务域下-合约数据
*/
func (m *EvfsBizContractRecord) Count(bizid string) int {
	count := 0
	db.Model(EvfsBizContractRecord{}).Where("biz_id=?",bizid).Count(&count)
	return count
}

//获取合约信息
func GetEvfsBizContractRecord() []EvfsBizContractRecord {
	data := make([]EvfsBizContractRecord, 0)
	db.Find(&data)
	return data
}

//规则申请同意
func (m *EvfsBizContractRecord) UpdateAddressCount(address string,count int64) error {
	return db.Model(&m).Where("contract_address", address).UpdateColumn("call_count", count).Error
}