package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "root:@tcp(mysql-container-restfull:3306)/demo"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Can not connect DB")
	}
	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
