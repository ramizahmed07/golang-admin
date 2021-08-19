package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	db "github.com/ramizahmed07/golang-admin/database"
	"github.com/ramizahmed07/golang-admin/routes"
)

func main() {
	db.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{AllowCredentials: true}))

	routes.Setup(app)

	app.Listen(":8000")
}
