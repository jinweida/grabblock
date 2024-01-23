package models

type EvfsStorageLicens struct {
	ReqId           string `gorm:"column:req_id;type:varchar(100);primary_key"` //代办ID
	TransactionHash string `gorm:"column:transaction_hash;type:varchar(100)"`   //更改规则交易hash 默认null
	CapacitySize    int64  `gorm:"column:capadity_size;"`                       //初始容量
	Y               int32  `gorm:"column:y"`                                    //年
	M               int32  `gorm:"column:m"`                                    //月
	D               int32  `gorm:"column:d"`                                    //日
	H               int32  `gorm:"column:h"`                                    //时
	AcceptTime      int64  `gorm:"column:accept_time;"`
	//AcceptTime      time.Time
	Approve         int32 `gorm:"column:approve;"` //审批状态 0=审核中 1=同意 2=拒绝

	Model
}
