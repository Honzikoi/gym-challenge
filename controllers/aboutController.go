package controllers

import (
	"strconv"

	"github.com/Honzikoi/gym-challenge/database"
	"github.com/Honzikoi/gym-challenge/models"
	"github.com/gofiber/fiber/v2"
)

// GetAbouts godoc
// @Summary Retrieve all abouts
// @Description Get all abouts
// @Tags abouts
// @Produce json
// @Success 200 {array} models.About
// @Router /abouts [get]
func GetAbouts(c *fiber.Ctx) error {
	var abouts []models.About
	database.DB.Db.Find(&abouts)
	return c.Status(fiber.StatusOK).JSON(abouts)
}

// CreateAbout godoc
// @Summary Create a new about
// @Description Create a new about
// @Tags abouts
// @Accept json
// @Produce json
// @Param about body models.About true "About"
// @Success 201 {object} models.About
// @Failure 400 {object} fiber.Map
// @Router /abouts [post]
func CreateAbout(c *fiber.Ctx) error {
	about := new(models.About)
	if err := c.BodyParser(about); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	database.DB.Db.Create(&about)
	return c.Status(fiber.StatusCreated).JSON(about)
}

// GetAbout godoc
// @Summary Retrieve an about by ID
// @Description Get an about by ID
// @Tags abouts
// @Produce json
// @Param id path int true "About ID"
// @Success 200 {object} models.About
// @Failure 400 {object} fiber.Map
// @Failure 404 {object} fiber.Map
// @Router /abouts/{id} [get]
func GetAbout(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid about ID"})
	}

	var about models.About
	result := database.DB.Db.First(&about, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "About not found"})
	}
	return c.Status(fiber.StatusOK).JSON(about)
}

// UpdateAbout godoc
// @Summary Update an about by ID
// @Description Update an about by ID
// @Tags abouts
// @Accept json
// @Produce json
// @Param id path int true "About ID"
// @Param about body models.About true "About"
// @Success 200 {object} models.About
// @Failure 400 {object} fiber.Map
// @Failure 404 {object} fiber.Map
// @Router /abouts/{id} [put]
func UpdateAbout(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid about ID"})
	}

	var about models.About
	result := database.DB.Db.First(&about, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "About not found"})
	}

	if err := c.BodyParser(&about); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	database.DB.Db.Save(&about)
	return c.Status(fiber.StatusOK).JSON(about)
}

// DeleteAbout godoc
// @Summary Delete an about by ID
// @Description Delete an about by ID
// @Tags abouts
// @Param id path int true "About ID"
// @Success 204
// @Failure 400 {object} fiber.Map
// @Failure 404 {object} fiber.Map
// @Router /abouts/{id} [delete]
func DeleteAbout(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid about ID"})
	}

	result := database.DB.Db.Delete(&models.About{}, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "About not found"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
