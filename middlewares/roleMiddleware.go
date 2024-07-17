package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

// Mock function to get user role from request
// In a real application, you would extract this from the JWT token or session
func getUserRole(c *fiber.Ctx) string {
    // This is a mock implementation. Replace this with your actual logic
    return c.Locals("userRole").(string)
}

func IsAdminOrCoach(c *fiber.Ctx) error {
    role := getUserRole(c)
    if role == "admin" || role == "coach" {
        return c.Next()
    }
    return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
        "error": "Access forbidden: requires admin or coach role",
    })
}
