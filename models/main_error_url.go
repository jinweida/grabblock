package models

import (
	"crypto/md5"
	"encoding/hex"

	"example.cn/grabblock/spider"
)

type MainErrorUrl struct {
	Hash      string `gorm:"column:hash;type:varchar(100);primary_key"`
	Url       string `gorm:"column:url;type:varchar(1000)"`
	FailCount int64  `gorm:"column:fail_count;"`
	Height    int64  `gorm:"column:height;"`
	TransHash string `gorm:"column:trans_hash;"`
	Model
}

func NewMainErrorUrl() *MainErrorUrl {
	return &MainErrorUrl{}
}
func (m *MainErrorUrl) Add(r *spider.Response) error {
	m.Hash = md5V(r.Request.Url)
	m.Url = r.Request.Url
	m.FailCount = r.Request.FailCount
	m.Height = r.Request.Height
	m.TransHash = r.Request.Hash

	return db.Create(&m).Error
}

func md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
