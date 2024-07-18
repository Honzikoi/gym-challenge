package controllers

import (
	"strconv"

	"github.com/Honzikoi/gym-challenge/database"
	"github.com/Honzikoi/gym-challenge/models"
	"github.com/gofiber/fiber/v2"
)

// GetGyms godoc
//
//	@Summary		Get all gyms
//	@Description	Get a list of all gyms
//	@Tags			Gym
//	@Produce		json
//	@Success		200	{array}	models.Gym
//	@Router			/gyms [get]
func GetGyms(c *fiber.Ctx) error {
	var gyms []models.Gym
	database.DB.Db.Preload("Equipment").Find(&gyms)
	return c.Status(fiber.StatusOK).JSON(gyms)
}

// GetGym godoc
//
//	@Summary		Get a gym by ID
//	@Description	Get a gym by its ID
//	@Tags			Gym
//	@Produce		json
//	@Param			id	path		int	true	"Gym ID"
//	@Success		200	{object}	models.Gym
//	@Failure		400	{string}	Invalid	gym	ID
//	@Failure		404	{string}	Gym	not	found
//	@Router			/gyms/{id} [get]
func GetGym(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid gym ID"})
	}

	var gym models.Gym
	if err := database.DB.Db.Preload("Equipment").First(&gym, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Gym not found"})
	}
	return c.Status(fiber.StatusOK).JSON(gym)
}

// CreateGym godoc
//
//	@Summary		Create a new gym
//	@Description	Create a new gym with name and location
//	@Tags			Gym
//	@Accept			json
//	@Produce		json
//	@Param			gym	body		models.Gym	true	"Gym"
//	@Success		201	{object}	models.Gym
//	@Failure		400	{object}	string: Invalid request payload
//	@Router			/gyms [post]
func CreateGym(c *fiber.Ctx) error {
	gym := new(models.Gym)
	if err := c.BodyParser(gym); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}
	database.DB.Db.Create(&gym)
	return c.Status(fiber.StatusCreated).JSON(gym)
}

// UpdateGym godoc
//
//	@Summary		Update a gym by ID
//	@Description	Update a gym's details by its ID
//	@Tags			Gym
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int			true	"Gym ID"
//	@Param			gym	body		models.Gym	true	"Gym"
//	@Success		200	{object}	models.Gym
//	@Failure		400	{string}	string	"Invalid gym ID"
//	@Failure		404	{string}	string	"Gym not found"
//	@Router			/gyms/{id} [put]
func UpdateGym(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid gym ID"})
	}

	var gym models.Gym
	if err := database.DB.Db.First(&gym, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Gym not found"})
	}

	if err := c.BodyParser(&gym); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	database.DB.Db.Save(&gym)
	return c.Status(fiber.StatusOK).JSON(gym)
}

// DeleteGym godoc
//
//	@Summary		Delete a gym by ID
//	@Description	Delete a gym by its ID
//	@Tags			Gym
//	@Param			id	path		int		true	"Gym ID"
//	@Success		204	{string}	string	"Successfully deleted"
//	@Failure		400	{string}	string	"Invalid gym ID"
//	@Failure		404	{string}	string	"Gym not found"
//	@Router			/gyms/{id} [delete]
func DeleteGym(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid gym ID"})
	}

	if err := database.DB.Db.Delete(&models.Gym{}, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Gym not found"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
