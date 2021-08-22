package database

import (
	"github.com/ramizahmed07/golang-admin/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "root:password@tcp(127.0.0.1:3306)/admindb?charset=utf8mb4&parseTime=True&loc=Local"

// DB - points to database
var DB *gorm.DB

// Connect - connects to database
func Connect() {
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Error: Could not connect to the database")
	}

	DB = database

	database.AutoMigrate(&models.User{}, &models.Role{})
}
