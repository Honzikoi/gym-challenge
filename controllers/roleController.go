package controllers

import (
	"strconv"

	"github.com/Honzikoi/gym-challenge/internal/database"
	"github.com/Honzikoi/gym-challenge/models"
	"github.com/gofiber/fiber/v2"
)

// GetRoles godoc
//
//	@Summary		Get all roles
//	@Description	Retrieve all roles
//	@Tags			Roles
//	@Produce		json
//	@Success		200	{array}		models.Role
//	@Failure		500	{string}	Internal Server Error
//	@Router			/roles [get]
func GetRoles(c *fiber.Ctx) error {
	var roles []models.Role
	if err := database.DB.Db.Find(&roles).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return c.Status(fiber.StatusOK).JSON(roles)
}

// CreateRole godoc
//
//	@Summary		Create a new role
//	@Description	Create a new role
//	@Tags			Roles
//	@Accept			json
//	@Produce		json
//	@Param			role	body		models.Role	true	"Role"
//	@Success		201	{object}	models.Role
//	@Failure		400	{string}	Cannot parse JSON
//	@Router			/roles [post]
func CreateRole(c *fiber.Ctx) error {
	role := new(models.Role)
	if err := c.BodyParser(role); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if err := database.DB.Db.Create(&role).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return c.Status(fiber.StatusCreated).JSON(role)
}

// GetRole godoc
//
//	@Summary		Get a role by ID
//	@Description	Retrieve a role by ID
//	@Tags			Roles
//	@Produce		json
//	@Param			id	path		int	true	"Role ID"
//	@Success		200	{object}	models.Role
//	@Failure		400	{string}	Invalid role ID
//	@Failure		404	{string}	Role not found
//	@Router			/roles/{id} [get]
func GetRole(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid role ID"})
	}

	var role models.Role
	result := database.DB.Db.First(&role, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Role not found"})
	}
	return c.Status(fiber.StatusOK).JSON(role)
}

// UpdateRole godoc
//
//	@Summary		Update a role by ID
//	@Description	Update a role's details by ID
//	@Tags			Roles
//	@Accept			json
//	@Produce		json
//	@Param			id		path	int			true	"Role ID"
//	@Param			role	body	models.Role	true	"Role"
//	@Success		200	{object}	models.Role
//	@Failure		400	{string}	Invalid role ID
//	@Failure		404	{string}	Role not found
//	@Router			/roles/{id} [put]
func UpdateRole(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid role ID"})
	}

	var role models.Role
	result := database.DB.Db.First(&role, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Role not found"})
	}

	if err := c.BodyParser(&role); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if err := database.DB.Db.Save(&role).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return c.Status(fiber.StatusOK).JSON(role)
}

// DeleteRole godoc
//
//	@Summary		Delete a role by ID
//	@Description	Delete a role by ID
//	@Tags			Roles
//	@Param			id	path	int	true	"Role ID"
//	@Success		204	{string}	string	"Successfully deleted"
//	@Failure		400	{string}	Invalid role ID
//	@Failure		404	{string}	Role not found
//	@Router			/roles/{id} [delete]
func DeleteRole(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid role ID"})
	}

	result := database.DB.Db.Delete(&models.Role{}, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Role not found"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
