package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ramizahmed07/golang-admin/database"
	"github.com/ramizahmed07/golang-admin/models"
)

// AllUsers - returns all users stored in database
func AllUsers(c *fiber.Ctx) error {
	var users []models.User

	database.DB.Find(&users)
	return c.JSON(users)
}

// CreateUser - creates user
func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	user.SetPassword("1234")

	database.DB.Create(&user)
	return c.JSON(user)
}

// GetUser - gets user by id
func GetUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	user := models.User{
		ID: uint(id),
	}
	database.DB.First(&user)
	return c.JSON(user)
}
