package handlers

import (
	"github.com/Honzikoi/gym-challenge/internal/database"
	"github.com/Honzikoi/gym-challenge/models"
	"github.com/gofiber/fiber/v2"
)

func AddGym(c *fiber.Ctx) error {
	type Request struct {
		Name     string `json:"name"`
		Location string `json:"location"`
	}

	var request Request
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	gym := models.Gym{
		Name:     request.Name,
		Location: request.Location,
	}

	database.DB.Db.Create(&gym)
	return c.JSON(gym)
}

func AddEquipmentToGym(c *fiber.Ctx) error {
	type Request struct {
		GymID         uint   `json:"gym_id"`
		EquipmentName string `json:"equipment_name"`
	}

	var request Request
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	var gym models.Gym
	if err := database.DB.Db.First(&gym, request.GymID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Gym not found"})
	}

	equipment := models.Equipment{
		Name: request.EquipmentName,
	}

	database.DB.Db.Model(&gym).Association("Equipment").Append(&equipment)
	return c.JSON(gym)
}
