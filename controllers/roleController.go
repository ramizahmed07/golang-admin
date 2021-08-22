package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ramizahmed07/golang-admin/database"
	"github.com/ramizahmed07/golang-admin/models"
)

// AllRoles - returns all roles stored in database
func AllRoles(c *fiber.Ctx) error {
	var roles []models.Role

	database.DB.Find(&roles)
	return c.JSON(roles)
}

// CreateRole - creates role
func CreateRole(c *fiber.Ctx) error {
	var role models.Role

	if err := c.BodyParser(&role); err != nil {
		return err
	}

	database.DB.Create(&role)

	return c.JSON(role)
}
