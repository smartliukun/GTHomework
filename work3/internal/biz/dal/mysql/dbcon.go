package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDB() (db *gorm.DB) {
	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	if db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		panic(err)
	}
	return
}
