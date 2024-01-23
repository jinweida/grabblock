package models

/***
  业务域-数据统计（业务域详情）
*/
type EvfsBizStatic struct {
	BizId              string `gorm:"column:biz_id;type:varchar(100);primary_key"`       //业务域ID
	BizName            string `gorm:"column:biz_name;type:varchar(100);"`                //业务域名称
	FileSize           int64 `gorm:"column:file_size;"`               //非结构化容量
	DataSize           int64 `gorm:"column:data_size;"`               //结构化容量
	OnlineFileSize     int64 `gorm:"column:online_file_size;"`        //上链非结构化容量
	OnlineFileCount    int64 `gorm:"column:online_file_count;"`       //上链文件条数
	OnlineDataSize     int64 `gorm:"column:online_data_size;"`        //上链结构化容量
	OnlineDataCount    int64 `gorm:"column:online_data_count;"`       //上链结构化条数
	ContractCount      int64 `gorm:"column:contract_count;"`          //合约数量
	OnlineUserCount    int64 `gorm:"column:online_user_count;"`       //上链用户数
	OrgCount           int64  `gorm:"column:org_count;"`              //企业个数
	SysCount           int64  `gorm:"column:sys_count;"`              //系统个数
	Model
}

//获取 统计 信息
func GetEvfsBizStatic() []EvfsBizStatic {
	data := make([]EvfsBizStatic, 0)
	db.Find(&data)
	return data
}


/**
  更新数据
*/
func (m *EvfsBizStatic) Save() error {
	return db.Save(&m).Error
}


/**
  分页查询 业务域统计信息
*/
func (m *EvfsBizStatic) Find(Page int,PageSize int) []EvfsBizStatic {
	data := make([]EvfsBizStatic, 0)
	if Page > 0 && PageSize > 0 {
		db.Order("created asc").Limit(PageSize).Offset((Page - 1) * PageSize).Find(&data)
	}else{
		db.Order("created asc").Limit(10).Offset(0).Find(&data)
	}
	return data
}

/**
  业务域统计数据
*/
func (m *EvfsBizStatic) Count() int {
	count := 0
	db.Model(EvfsBizStatic{}).Count(&count)
	return count
}


func (m *EvfsBizStatic) Add() error {
	return db.Create(&m).Error
}


