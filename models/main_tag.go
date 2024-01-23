package models

type MainTag struct {
	Address string `gorm:"column:address;type:varchar(100);primary_key"`
	Tag     string `gorm:"column:tag;type:varchar(50)"`
	Status  int32  `gorm:"column:status;"`
	Model
}
