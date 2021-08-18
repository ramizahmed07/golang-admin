package main

import (
	"github.com/gofiber/fiber/v2"
	db "github.com/ramizahmed07/golang-admin/database"
)

func main() {
	db.Connect()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":8000")
}
