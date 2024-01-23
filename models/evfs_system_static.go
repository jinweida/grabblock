package models

/***
  业务域-系统信息统计（业务域下-某系统详情）
*/
type EvfsSystemStatic struct {
	SysId              string `gorm:"column:sys_id;type:varchar(100);primary_key"`    //系统ID
	SysName            string `gorm:"column:sys_name;type:varchar(100);"`             //系统名称
	BizId              string `gorm:"column:biz_id;type:varchar(100);"`               //业务域ID
	OrgId              string `gorm:"column:org_id;type:varchar(100);"`               //企业ID
	OrgName            string `gorm:"column:org_name;type:varchar(100);"`             //企业名称
	OnlineFileSize     int64 `gorm:"column:online_file_size;"`                          //上链非结构化容量
	OnlineFileCount    int64 `gorm:"column:online_file_count;"`                         //上链文件条数
	OnlineDataSize     int64 `gorm:"column:online_data_size;"`                          //上链结构化容量
	OnlineDataCount    int64 `gorm:"column:online_data_count;"`                         //上链结构化条数
	OnlineUserCount    int64 `gorm:"column:online_user_count;"`                         //上链用户数
	Model
}

/**
  分页查询 某业务域下数据
 */
func (m *EvfsSystemStatic) Find(bizid string,Page int,PageSize int) []EvfsSystemStatic {
	data := make([]EvfsSystemStatic, 0)
	if Page > 0 && PageSize > 0 {
		db.Order("created asc").Where("biz_id=?",bizid).Limit(PageSize).Offset((Page - 1) * PageSize).Find(&data)
	}else{
		db.Order("created asc").Where("biz_id=?",bizid).Limit(10).Offset(0).Find(&data)
	}
	return data
}

/**
   某业务域下数据
 */
func (m *EvfsSystemStatic) Count(bizid string) int {
	count := 0
	db.Model(EvfsSystemStatic{}).Where("biz_id=?",bizid).Count(&count)
	return count
}


//获取合约信息
func GetEvfsSystemStatic() []EvfsSystemStatic {
	data := make([]EvfsSystemStatic, 0)
	db.Find(&data)
	return data
}

/**
  更新数据
*/
func (m *EvfsSystemStatic) Save() error {
	return db.Save(&m).Error
}

//新增数据
func (m *EvfsSystemStatic) Add() error {
	return db.Create(&m).Error
}