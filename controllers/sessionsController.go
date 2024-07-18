package controllers

import (
	"strconv"

	"github.com/Honzikoi/gym-challenge/database"
	"github.com/Honzikoi/gym-challenge/models"
	"github.com/gofiber/fiber/v2"
)

// GetSessions godoc
//	@Summary		Retrieve all sessions
//	@Description	Get all sessions
//	@Tags			sessions
//	@Produce		json
//	@Success		200	{array}	models.Sessions
//	@Router			/sessions [get]
func GetSessions(c *fiber.Ctx) error {
	var sessions []models.Sessions
	database.DB.Db.Find(&sessions)
	return c.Status(fiber.StatusOK).JSON(sessions)
}

// CreateSession godoc
//	@Summary		Create a new session
//	@Description	Create a new session
//	@Tags			sessions
//	@Accept			json
//	@Produce		json
//	@Param			session	body		models.Sessions	true	"Session"
//	@Success		201		{object}	models.Sessions
//	@Failure		400		{object}	fiber.Map
//	@Router			/sessions [post]
func CreateSession(c *fiber.Ctx) error {
	session := new(models.Sessions)
	if err := c.BodyParser(session); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	database.DB.Db.Create(&session)
	return c.Status(fiber.StatusCreated).JSON(session)
}

// GetSession godoc
//	@Summary		Retrieve a session by ID
//	@Description	Get a session by ID
//	@Tags			sessions
//	@Produce		json
//	@Param			id	path		int	true	"Session ID"
//	@Success		200	{object}	models.Sessions
//	@Failure		400	{object}	fiber.Map
//	@Failure		404	{object}	fiber.Map
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
//	@Summary		Update a session by ID
//	@Description	Update a session by ID
//	@Tags			sessions
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int				true	"Session ID"
//	@Param			session	body		models.Sessions	true	"Session"
//	@Success		200		{object}	models.Sessions
//	@Failure		400		{object}	fiber.Map
//	@Failure		404		{object}	fiber.Map
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

	database.DB.Db.Save(&session)
	return c.Status(fiber.StatusOK).JSON(session)
}

// DeleteSession godoc
//	@Summary		Delete a session by ID
//	@Description	Delete a session by ID
//	@Tags			sessions
//	@Param			id	path	int	true	"Session ID"
//	@Success		204
//	@Failure		400	{object}	fiber.Map
//	@Failure		404	{object}	fiber.Map
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
