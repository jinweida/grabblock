package models

type EvfsStorageMember struct {
	MemberId        int32     `gorm:"column:member_id;primary_key"`
	TransactionHash string    `gorm:"column:transaction_hash;type:varchar(100);"` //交易hash
	ReqId           string    `gorm:"column:req_id;type:varchar(100);"` //待办ID
	AddressName     string    `gorm:"column:address_name;type:varchar(100);"`//存管域--管理员名称
	Address         string    `gorm:"column:address;type:varchar(100);"` //存管域--管理员地址
	DomainId        string    `gorm:"column:domain_id;type:varchar(100);"` //存储域ID
	Name            string    `gorm:"column:name;type:varchar(100);"` //存管域名称
	JoinTime        int64      `gorm:"column:jsoin_time;"` //加入日期
	Approve         int32     `gorm:"column:approve;"`                 //审批状态 0=审核中 1=同意 2=拒绝
	Op              int64     `gorm:"column:op;"`                 //操作 1 添加 2移除
	Model
}

//成员申请
func (m *EvfsStorageMember) Add() error {
	return db.Create(&m).Error
}

//成员同意
func (m *EvfsStorageMember) Delete() error {
	return db.Model(&m).Where("domain_id=? and address=?", m.DomainId,m.Address).Delete(m).Error
}

//成员同意
func (m *EvfsStorageMember) UpdateApprove() error {
	//如果数据库已有此数据,则认为此数据是删除同意，将删除使用数据
	data := make([]EvfsStorageMember, 0)
	db.Where("req_id=?",m.ReqId).Find(&data)
	if len(data)>0 && data[0].Op==2 {
		//删除同意
		return db.Model(&m).Where("domain_id=? and address=?", m.DomainId,data[0].Address).Delete(m).Error
	}else{
		//添加同意
		return db.Model(&m).Where("req_id=?", m.ReqId).UpdateColumn("approve", m.Approve).Error
	}
}
