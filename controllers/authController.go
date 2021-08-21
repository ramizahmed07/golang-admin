package controllers

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ramizahmed07/golang-admin/database"
	"github.com/ramizahmed07/golang-admin/models"
	utils "github.com/ramizahmed07/golang-admin/utils"
)

// Register - home route controller
func Register(c *fiber.Ctx) error {
	var payload = map[string]string{}
	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	if payload["password"] != payload["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Passwords don't match",
		})
	}

	user := models.User{
		FirstName: payload["first_name"],
		LastName:  payload["last_name"],
		Email:     payload["email"],
	}
	user.SetPassword(payload["password"])
	database.DB.Create(&user)

	return c.JSON(user)
}

// Login - logins user
func Login(c *fiber.Ctx) error {
	var payload map[string]string

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("email = ?", payload["email"]).First(&user)

	if user.ID == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "User not found",
		})
	}

	if err := user.ComparePassword(payload["password"]); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Password is incorrect",
		})
	}

	token, err := utils.GenerateJwt(strconv.Itoa(int(user.ID)))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	//create cookie
	cookie := new(fiber.Cookie)
	cookie.Name = "jwt"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HTTPOnly = true
	// Set cookie
	c.Cookie(cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

// User - gets logged in user
func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	id, _ := utils.ParseJwt(cookie)

	var user models.User
	database.DB.Where("id = ?", id).First(&user)
	return c.JSON(user)
}

//Logout - logout user
func Logout(c *fiber.Ctx) error {
	cookie := new(fiber.Cookie)
	cookie.Name = "jwt"
	cookie.Value = ""
	cookie.Expires = time.Now().Add(-time.Hour)
	cookie.HTTPOnly = true
	c.Cookie(cookie)
	return c.JSON(fiber.Map{
		"message": "success",
	})

}
