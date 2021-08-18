package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "root:password@tcp(127.0.0.1:3306)/admindb?charset=utf8mb4&parseTime=True&loc=Local"

func main() {
	app := fiber.New()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error: Could not connect to the database")
	}
	fmt.Println(db)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":8000")
}
