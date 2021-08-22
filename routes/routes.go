package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ramizahmed07/golang-admin/controllers"
	"github.com/ramizahmed07/golang-admin/middlewares"
)

// Setup routes
func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Use(middlewares.IsAuthenticated)
	app.Get("/api/user", controllers.User)
	app.Get("/api/logout", controllers.Logout)
	app.Get("/api/users", controllers.AllUsers)
	app.Post("/api/users", controllers.CreateUser)
	app.Get("/api/users/:id", controllers.GetUser)
	app.Put("/api/users/:id", controllers.UpdateUser)
	app.Delete("/api/users/:id", controllers.DeleteUser)
	app.Get("/api/roles", controllers.AllRoles)
	app.Get("/api/roles/:id", controllers.GetRole)
	app.Post("/api/roles", controllers.CreateRole)
	app.Put("/api/roles/:id", controllers.UpdateRole)
	app.Delete("/api/roles/:id", controllers.DeleteRole)

}
