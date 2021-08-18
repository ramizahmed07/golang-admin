package database

import (
	"github.com/ramizahmed07/golang-admin/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "root:password@tcp(127.0.0.1:3306)/admindb?charset=utf8mb4&parseTime=True&loc=Local"

// Connect - connects to database
func Connect() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Error: Could not connect to the database")
	}

	db.AutoMigrate(&models.User{})
}
