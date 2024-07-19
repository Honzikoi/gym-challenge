package handlers

import (
	"github.com/Honzikoi/gym-challenge/internal/database"
	"github.com/Honzikoi/gym-challenge/models"
	"github.com/gofiber/fiber/v2"
)

func AssignRole(c *fiber.Ctx) error {
	type Request struct {
		UserID uint `json:"user_id"`
		RoleID uint `json:"role_id"`
	}

	var request Request
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	var user models.Users
	var role models.Role

	if err := database.DB.Db.First(&user, request.UserID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}
	if err := database.DB.Db.First(&role, request.RoleID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Role not found"})
	}

	database.DB.Db.Model(&user).Association("Roles").Append(&role)
	return c.JSON(user)
}
