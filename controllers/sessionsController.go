package controllers

import (
	"strconv"

	"github.com/Honzikoi/gym-challenge/database"
	"github.com/Honzikoi/gym-challenge/models"
	"github.com/gofiber/fiber/v2"
)

// GetSessions - Retrieve all sessions
func GetSessions(c *fiber.Ctx) error {
    var sessions []models.Sessions
    database.DB.Db.Find(&sessions)
    return c.Status(fiber.StatusOK).JSON(sessions)
}

// CreateSession - Create a new session
func CreateSession(c *fiber.Ctx) error {
    session := new(models.Sessions)
    if err := c.BodyParser(session); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
    }
    database.DB.Db.Create(&session)
    return c.Status(fiber.StatusCreated).JSON(session)
}

// GetSession - Retrieve a session by ID
func GetSession(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid session ID"})
    }

    var session models.Sessions
    result := database.DB.Db.First(&session, id)
    if result.Error != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Session not found"})
    }
    return c.Status(fiber.StatusOK).JSON(session)
}

// UpdateSession - Update a session by ID
func UpdateSession(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid session ID"})
    }

    var session models.Sessions
    result := database.DB.Db.First(&session, id)
    if result.Error != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Session not found"})
    }

    if err := c.BodyParser(&session); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
    }

    database.DB.Db.Save(&session)
    return c.Status(fiber.StatusOK).JSON(session)
}

// DeleteSession - Delete a session by ID
func DeleteSession(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid session ID"})
    }

    result := database.DB.Db.Delete(&models.Sessions{}, id)
    if result.Error != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Session not found"})
    }

    return c.SendStatus(fiber.StatusNoContent)
}
