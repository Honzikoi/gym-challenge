package controllers

import (
	"strconv"

	"github.com/Honzikoi/gym-challenge/internal/database"
	"github.com/Honzikoi/gym-challenge/models"
	"github.com/gofiber/fiber/v2"
)

// GetGroups godoc
//
//	@Summary		Get all groups
//	@Description	Retrieve all groups
//	@Tags			Groups
//	@Produce		json
//	@Success		200	{array}		models.Group
//	@Failure		500	{string}	Internal Server Error
//	@Router			/groups [get]
func GetGroups(c *fiber.Ctx) error {
	var groups []models.Group
	if err := database.DB.Db.Find(&groups).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return c.Status(fiber.StatusOK).JSON(groups)
}

// CreateGroup godoc
//
//	@Summary		Create a new group
//	@Description	Create a new group
//	@Tags			Groups
//	@Accept			json
//	@Produce		json
//	@Param			group	body		models.Group	true	"Group"
//	@Success		201	{object}	models.Group
//	@Failure		400	{string}	Cannot parse JSON
//	@Failure		500	{string}	Internal Server Error
//	@Router			/groups [post]
func CreateGroup(c *fiber.Ctx) error {
	group := new(models.Group)
	if err := c.BodyParser(group); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if err := database.DB.Db.Create(&group).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return c.Status(fiber.StatusCreated).JSON(group)
}

// GetGroup godoc
//
//	@Summary		Get a group by ID
//	@Description	Retrieve a group by ID
//	@Tags			Groups
//	@Produce		json
//	@Param			id	path		int	true	"Group ID"
//	@Success		200	{object}	models.Group
//	@Failure		400	{string}	Invalid group ID
//	@Failure		404	{string}	Group not found
//	@Router			/groups/{id} [get]
func GetGroup(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid group ID"})
	}

	var group models.Group
	result := database.DB.Db.First(&group, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Group not found"})
	}
	return c.Status(fiber.StatusOK).JSON(group)
}

// UpdateGroup godoc
//
//	@Summary		Update a group by ID
//	@Description	Update a group's details by ID
//	@Tags			Groups
//	@Accept			json
//	@Produce		json
//	@Param			id		path	int			true	"Group ID"
//	@Param			group	body	models.Group	true	"Group"
//	@Success		200	{object}	models.Group
//	@Failure		400	{string}	Invalid group ID
//	@Failure		404	{string}	Group not found
//	@Router			/groups/{id} [put]
func UpdateGroup(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid group ID"})
	}

	var group models.Group
	result := database.DB.Db.First(&group, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Group not found"})
	}

	if err := c.BodyParser(&group); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if err := database.DB.Db.Save(&group).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return c.Status(fiber.StatusOK).JSON(group)
}

// DeleteGroup godoc
//
//	@Summary		Delete a group by ID
//	@Description	Delete a group by ID
//	@Tags			Groups
//	@Param			id	path	int	true	"Group ID"
//	@Success		204	{string}	Successfully deleted
//	@Failure		400	{string}	Invalid group ID
//	@Failure		404	{string}	Group not found
//	@Router			/groups/{id} [delete]
func DeleteGroup(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid group ID"})
	}

	result := database.DB.Db.Delete(&models.Group{}, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Group not found"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
