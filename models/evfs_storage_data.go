package models

import (
	"fclink.cn/grabblock/tools/log"
	"time"
)

type EvfsStorageData struct {
	DataId          string    `gorm:"column:data_id;type:varchar(100);primary_key"`
	DataSize        int64     `gorm:"column:data_size;"`
	BizId           string    `gorm:"column:biz_id;type:varchar(100);"`
	SystemId        string    `gorm:"column:system_id;type:varchar(100);"`
	ResourceId      string    `gorm:"column:resource_id;type:varchar(100);"`
	Y               int32     `gorm:"column:y"` //年
	M               int32     `gorm:"column:m"` //月
	D               int32     `gorm:"column:d"` //日
	H               int32     `gorm:"column:h"` //时
	TransactionHash string    `gorm:"column:transaction_hash;type:varchar(100)"`
	DataOwner        string    `gorm:"column:data_owner;type:varchar(100)"`
	Op               int32      `gorm:"column:op"` //操作 1 添加 2修改
	MarkDelete       int32      `gorm:"column:mark_delete"` //操作 0 否 1是
	DataHash         string    `gorm:"column:data_hash;type:varchar(100);"`//文件hash
	DeleteTime       int64    `gorm:"column:delete_time;"`//删除时间戳
	UplineTime       int64    `gorm:"column:upline_time;"`//完成上链时间
	ApplyTime        int64    `gorm:"column:apply_time;"`//申请上链时间戳
	Approve          int32     `gorm:"column:approve;"`                    //审批状态 0=审核中 1=同意 2=拒绝
	Model
}

/**
  申请上传结构化数据
*/
func (t *EvfsStorageData) ApplyData() error {
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

	if err := tx.Create(&t).Error; err != nil {
		tx.Rollback()
		log.Errorf("err:%s", err.Error())
		return err
	}
	return tx.Commit().Error
}

/**
  结构化数据上链确认删除
*/
func (t *EvfsStorageData) DataDeleteConfirm(datahash string,deletetime int64) error {
	return db.Model(&t).Where("data_hash=?", datahash).UpdateColumn("mark_delete", 1).UpdateColumn("delete_time", deletetime).Error
}



/**
  统计业务域下文件容量、个数数据
*/
func (m *EvfsStorageData) StaticSizeAndCount(bizid string) (int64,int64) {
	//SELECT sum(data_size)data_size,count(*)apply_time FROM evfs_storage_data where biz_id=1
	db.Model(m).Select("sum(data_size)data_size,count(*)apply_time").Where("biz_id=? and approve=1",bizid).Find(&m)
	return m.DataSize,m.ApplyTime
}

/**
  统计业务域下文件容量、个数数据
*/
func (m *EvfsStorageData) StaticSizeAndCountInsys(bizid string,sysid string) (int64,int64) {
	//SELECT sum(data_size)data_size,count(*)apply_time FROM evfs_storage_data where biz_id=1
	db.Model(m).Select("sum(data_size)data_size,count(*)apply_time").Where("biz_id=? and system_id=? and approve=1",bizid,sysid).Find(&m)
	return m.DataSize,m.ApplyTime
}

/**
  24小时内上传结构化数据条数
*/
func (m *EvfsStorageData) UpDataIn24Count() int64 {
	//SELECT count(*)apply_time FROM evfs_storage_data where upline_time BETWEEN 1504476474 and 1604476474
	timeinfo := time.Now()
	endtime := timeinfo.Unix()
	starttime := endtime-(24*60*60)
	db.Model(m).Select("count(*)apply_time").Where("upline_time BETWEEN ? and ?",starttime,endtime).Find(&m)
	return m.ApplyTime
}