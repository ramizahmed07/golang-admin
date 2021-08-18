package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ramizahmed07/golang-admin/models"
)

// Register - home route controller
func Register(c *fiber.Ctx) error {
	user := models.User{
		ID:        1,
		FirstName: "Ramizzz",
		LastName:  "Ahmed",
		Email:     "ramizahmedar@gmail.com",
		Password:  "ramiz123",
	}
	return c.JSON(user)
	// return c.SendString("Hello, World 👋!")
}
