package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ramizahmed07/golang-admin/models"
	"golang.org/x/crypto/bcrypt"
)

// Register - home route controller
func Register(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "passwords don't match",
		})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
		Password:  password,
	}
	return c.JSON(user)
	// return c.SendString("Hello, World 👋!")
}
