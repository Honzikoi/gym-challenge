package controllers

import (
	"strconv"

	"github.com/Honzikoi/gym-challenge/database"
	"github.com/Honzikoi/gym-challenge/models"
	"github.com/gofiber/fiber/v2"
)

// GetAbouts - Retrieve all abouts
func GetAbouts(c *fiber.Ctx) error {
    var abouts []models.About
    database.DB.Db.Find(&abouts)
    return c.Status(fiber.StatusOK).JSON(abouts)
}

// CreateAbout - Create a new about
func CreateAbout(c *fiber.Ctx) error {
    about := new(models.About)
    if err := c.BodyParser(about); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
    }
    database.DB.Db.Create(&about)
    return c.Status(fiber.StatusCreated).JSON(about)
}

// GetAbout - Retrieve an about by ID
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

// UpdateAbout - Update an about by ID
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

// DeleteAbout - Delete an about by ID
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
