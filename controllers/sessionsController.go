package controllers

import (
	"strconv"
	"time"

	"github.com/Honzikoi/gym-challenge/database"
	"github.com/Honzikoi/gym-challenge/models"
	"github.com/gofiber/fiber/v2"
)

// GetSessions godoc
//
//	@Summary		Get all sessions
//	@Description	Retrieve all sessions
//	@Tags			Sessions
//	@Produce		json
//	@Success		200	{array}		models.Sessions
//	@Failure		500	{string}	Internal Server Error
//	@Router			/sessions [get]
func GetSessions(c *fiber.Ctx) error {
	var sessions []models.Sessions
	if err := database.DB.Db.Find(&sessions).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return c.Status(fiber.StatusOK).JSON(sessions)
}

// CreateSession godoc
//
//	@Summary		Create a new session
//	@Description	Create a new session
//	@Tags			Sessions
//	@Accept			json
//	@Produce		json
//	@Param			session	body		models.Sessions	true	"Session"
//	@Success		201	{object}	models.Sessions
//	@Failure		400	{string}	Cannot parse JSON
//	@Router			/sessions [post]
func CreateSession(c *fiber.Ctx) error {
	sessions := new(models.Sessions)
	if err := c.BodyParser(sessions); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// Parse the date field
	if sessions.Date.IsZero() {
		sessions.Date = time.Now() // set current time if not provided
	}

	if err := database.DB.Db.Create(&sessions).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return c.Status(fiber.StatusCreated).JSON(sessions)
}

// GetSession godoc
//
//	@Summary		Get a session by ID
//	@Description	Retrieve a session by ID
//	@Tags			Sessions
//	@Produce		json
//	@Param			id	path		int	true	"Session ID"
//	@Success		200	{object}	models.Sessions
//	@Failure		400	{string}	Invalid session ID
//	@Failure		404	{string}	Session not found
//	@Router			/sessions/{id} [get]
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

// UpdateSession godoc
//
//	@Summary		Update a session by ID
//	@Description	Update a session's details by ID
//	@Tags			Sessions
//	@Accept			json
//	@Produce		json
//	@Param			id		path	int				true	"Session ID"
//	@Param			session	body	models.Sessions	true	"Session"
//	@Success		200	{object}	models.Sessions
//	@Failure		400	{string}	Invalid session ID
//	@Failure		404	{string}	Session not found
//	@Router			/sessions/{id} [put]
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

	if err := database.DB.Db.Save(&session).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return c.Status(fiber.StatusOK).JSON(session)
}

// DeleteSession godoc
//
//	@Summary		Delete a session by ID
//	@Description	Delete a session by ID
//	@Tags			Sessions
//	@Param			id	path	int	true	"Session ID"
//	@Success		204	{string}	Successfully deleted
//	@Failure		400	{string}	Invalid session ID
//	@Failure		404	{string}	Session not found
//	@Router			/sessions/{id} [delete]
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
