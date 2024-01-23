package models

import (
	"fclink.cn/grabblock/tools/log"
	"time"
)

type EvfsStorageFile struct {
	FileId          string    `gorm:"column:file_id;type:varchar(100);primary_key"`
	FileName        string    `gorm:"column:file_name;type:varchar(500);"`
	FileSize        int64     `gorm:"column:file_size;"`
	BizId           string    `gorm:"column:biz_id;type:varchar(100);"`
	SystemId        string    `gorm:"column:system_id;type:varchar(100);"`
	ResourceId      string    `gorm:"column:resource_id;type:varchar(100);"`
	FileType        string    `gorm:"column:file_type;type:varchar(20);"`
	ExtType         int32     `gorm:"column:ext_type;"`
	CopyCount       int32     `gorm:"column:copy_count;"`
	SliceCount      int32     `gorm:"column:slice_count;"`
	Y               int32     `gorm:"column:y"` //年
	M               int32     `gorm:"column:m"` //月
	D               int32     `gorm:"column:d"` //日
	H               int32     `gorm:"column:h"` //时
	TransactionHash  string    `gorm:"column:transaction_hash;type:varchar(100)"`
	SendFrontAddress string    `gorm:"column:send_front_address;type:varchar(100)"`
	FileOwner        string    `gorm:"column:file_owner;type:varchar(100)"`
	Op               int32      `gorm:"column:op"` //操作 1 添加 2修改
	MarkDelete       int32      `gorm:"column:mark_delete"` //操作 0 否 1是
	FileHash         string    `gorm:"column:file_hash;type:varchar(100);"`//文件hash
	DeleteTime       int64    `gorm:"column:delete_time;"`//删除时间戳
	UplineTime       int64    `gorm:"column:upline_time;"`//完成上链时间
	ApplyTime        int64    `gorm:"column:apply_time;"`//申请上链时间戳
	Approve          int32     `gorm:"column:approve;"`                    //审批状态 0=审核中 1=同意 2=拒绝
	Model
}

/**
  申请上传非结构化数据
 */
func (t *EvfsStorageFile) ApplyFile() error {
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
   非结构化数据上链确认
   每次文件上传或修改的交易hash是文件的版本hash. 所以确认上传时，根据版本hash确认某文件是否上传成功
 */
func (t *EvfsStorageFile) ConfirmFile(txhash string,confirmtime int64) error {
	return db.Model(&t).Where("transaction_hash=?", txhash).UpdateColumn("upline_time", confirmtime).UpdateColumn("approve", 1).Error
}


/**
  非结构化数据上链确认删除
*/
func (t *EvfsStorageFile) FileDeleteConfirmData(filehash string,deletetime int64) error {
	return db.Model(&t).Where("file_hash=?", filehash).UpdateColumn("mark_delete", 1).UpdateColumn("delete_time", deletetime).Error
}


/**
  统计业务域下文件容量、个数数据
*/
func (m *EvfsStorageFile) StaticSizeAndCount(bizid string) (int64,int64) {
	//SELECT sum(file_size)file_size,count(*)apply_time FROM evfs_storage_file where biz_id=1
	db.Model(EvfsStorageFile{}).Select("sum(file_size)file_size,count(*)apply_time").Where("biz_id=? and approve=1",bizid).Find(&m)
	return m.FileSize,m.ApplyTime
}

/**
  统计业务域下文件容量、个数数据
*/
func (m *EvfsStorageFile) StaticSizeAndCountInsys(bizid string,sysid string) (int64,int64) {
	//SELECT sum(file_size)file_size,count(*)apply_time FROM evfs_storage_file where biz_id=1
	db.Model(EvfsStorageFile{}).Select("sum(file_size)file_size,count(*)apply_time").Where("biz_id=? and system_id=? and approve=1",bizid,sysid).Find(&m)
	return m.FileSize,m.ApplyTime
}


/**
  24小时内上传文件个数
*/
func (m *EvfsStorageFile) UpFileIn24Count() int64 {
	//SELECT count(*)apply_time FROM evfs_storage_file where upline_time BETWEEN 1504476474 and 1704476474
	timeinfo := time.Now()
	endtime := timeinfo.Unix()
	starttime := endtime-(24*60*60)
	db.Model(m).Select("count(*)apply_time").Where("upline_time BETWEEN ? and ?",starttime,endtime).Find(&m)
	return m.ApplyTime
}


