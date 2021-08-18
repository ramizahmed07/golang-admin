package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ramizahmed07/golang-admin/controllers"
)

// Setup routes
func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
}
