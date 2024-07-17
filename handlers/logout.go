// handlers/logout.go
package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// Logout handles the user logout
func Logout(c *fiber.Ctx) error {
	// Clear the JWT token cookie
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	})

	return c.JSON(fiber.Map{
		"message": "Successfully logged out",
	})
}
