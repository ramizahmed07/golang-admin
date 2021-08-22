package controllers

import (
	"strconv"

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

// GetRole - gets user by id
func GetRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	role := models.Role{
		ID: uint(id),
	}
	database.DB.First(&role)
	return c.JSON(role)
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

// UpdateRole - updates role
func UpdateRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		ID: uint(id),
	}

	if err := c.BodyParser(&role); err != nil {
		return err
	}

	database.DB.Model(&role).Updates(role)

	return c.JSON(role)
}
