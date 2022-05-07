package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	db *gorm.DB
)

func InitDB() {
	mydb, err := gorm.Open(mysql.Open("root:123456@(127.0.0.1:3306)/adx?charset=utf8mb4&parseTime=True&loc=Local"))
	db = mydb
	if err != nil {
		log.Fatalf("db init failed")
	}
}

func GetDB() *gorm.DB {
	return db
}
