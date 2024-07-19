package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// Logout godoc
//	@Summary		Log out a user
//	@Description	Invalidate the user's session or token
//	@Tags			Signup & Login
//	@Produce		json
//	@Success		200	{string}	Successfully	logged	out
//	@Router			/logout [post]
func Logout(c *fiber.Ctx) error {
	// Invalidate the user's session or token here
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully logged out",
	})
}
