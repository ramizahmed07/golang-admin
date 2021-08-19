package middlewares

import (
	"github.com/gofiber/fiber/v2"
	utils "github.com/ramizahmed07/golang-admin/utils"
)

// IsAuthenticated - auth middleware
func IsAuthenticated(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	if _, err := utils.ParseJwt(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	return c.Next()
}
