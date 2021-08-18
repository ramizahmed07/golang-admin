package controllers

import "github.com/gofiber/fiber/v2"

// Hello - home route controller
func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}
