package handlers

import (
	"github.com/Honzikoi/gym-challenge/database"
	"github.com/Honzikoi/gym-challenge/models"
	"github.com/gofiber/fiber/v2"
)

func RequestJoinGroup(c *fiber.Ctx) error {
	type Request struct {
		GroupID uint `json:"group_id"`
		UserID  uint `json:"user_id"`
	}

	var request Request
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// Check if the group exists
	var group models.Group
	if err := database.DB.Db.First(&group, request.GroupID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Group not found"})
	}

	// Create a new entry in the group_users table to associate the user with the group
	userGroup := models.Group{
		UserID:  request.UserID,
		GroupID: request.GroupID,
		Status:  "pending",
	}

	if err := database.DB.Db.Create(&userGroup).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to join group"})
	}

	return c.JSON(userGroup)
}
