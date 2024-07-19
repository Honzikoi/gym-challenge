package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func IsAdminOrCoach(c *fiber.Ctx) error {
	role := c.Locals("role")
	if role == "admin" || role == "coach" {
		return c.Next()
	}
	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
		"error": "Access denied",
	})
}
