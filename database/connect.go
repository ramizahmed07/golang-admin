package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "root:password@tcp(127.0.0.1:3306)/admindb?charset=utf8mb4&parseTime=True&loc=Local"

func Connect() {
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error: Could not connect to the database")
	}
}
