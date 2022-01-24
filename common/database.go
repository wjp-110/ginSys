package common

import (
	"ginSys/model"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:123456@/gin_sys?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database, err---->" + err.Error())
	}
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
