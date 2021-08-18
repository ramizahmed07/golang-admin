package main

import (
	"github.com/gofiber/fiber/v2"
	db "github.com/ramizahmed07/golang-admin/database"
	"github.com/ramizahmed07/golang-admin/routes"
)

func main() {
	db.Connect()
	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8000")
}
