package config

import (
	// "github.com/jinzhu/gorm"     //mysql
	// _ "github.com/jinzhu/gorm/dialects/mysql"   //mysql
	"gorm.io/driver/postgres" // postgres
	"gorm.io/gorm"            //postgres
)

var (
	db *gorm.DB
)

func Connect() {
	// d, err := gorm.Open("mysql", "root:Admin@123*#()$@/bookinfo?charset=utf8&parseTime=True&loc=Local")    // mysql

	dsn := "host=localhost user=postgres password=Admin@123*#()$ dbname=bookinfo port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
