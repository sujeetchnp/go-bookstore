package config 

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db * gorm.DB
)
	

func Connect(){
	d, err := gorm.Open("mysql", "sujeet:Admin@123@/simplerest?charset=utf8&parseTime=TYrue&loc=Local")
	err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB{
	return db
}