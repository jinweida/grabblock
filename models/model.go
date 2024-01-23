package models

import (
	"fmt"
	"log"
	"time"

	"example.cn/grabblock/conf"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Model struct {
	CreatedAt time.Time `gorm:"column:created" json:"created"`
	UpdatedAt time.Time `gorm:"column:updated" json:"updated"`
}

var db *gorm.DB

var org map[string]*EvfsOrg

func GetDB() *gorm.DB {
	return db
}

// db init
func Setup() {
	var err error
	c := conf.Context.Db

	args := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", c.Username, c.Password, c.Address, c.Port, c.Dbname)

	db, err = gorm.Open("mysql", args)

	if err != nil {
		log.Fatalln(err)
	}
	db.DB().SetMaxIdleConns(c.Maxidleconns)
	db.DB().SetMaxOpenConns(c.Maxopenconns)
	db.DB().SetConnMaxLifetime(c.Maxlifetime * time.Minute)
	db.DB().Ping()
	db.SingularTable(true)
	db.LogMode(c.LogMode)

	//db.CreateTable(&MainContract{}, &MainToken{})

	//db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Test{})
	if conf.Context.Node.Contractparsing {
		org = make(map[string]*EvfsOrg, 0)
		data := make([]EvfsOrg, 0)
		db.Find(&data)

		for _, m := range data {
			org[m.OrgId] = &m
		}
	}
}
