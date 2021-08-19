package controllers

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/ramizahmed07/golang-admin/database"
	"github.com/ramizahmed07/golang-admin/models"
	"golang.org/x/crypto/bcrypt"
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

	password, _ := bcrypt.GenerateFromPassword([]byte(payload["password"]), 14)

	user := models.User{
		FirstName: payload["first_name"],
		LastName:  payload["last_name"],
		Email:     payload["email"],
		Password:  password,
	}
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

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(payload["password"])); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Password is incorrect",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 1 day
	})

	token, err := claims.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(token)
}
