package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func CheckRole(role string) fiber.Handler {
    return func(c *fiber.Ctx) error {
        userRoles := c.Locals("roles").([]string)
        isAdmin := false
        hasRole := false
        for _, r := range userRoles {
            if r == "admin" {
                isAdmin = true
            }
            if r == role {
                hasRole = true
            }
        }
        if !isAdmin && !hasRole {
            return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Permission denied"})
        }
        return c.Next()
    }
}
