package controllers

import (
	"strconv"

	"github.com/Honzikoi/gym-challenge/database"
	"github.com/Honzikoi/gym-challenge/models"
	"github.com/gofiber/fiber/v2"
)

// GetUsers - Retrieve all users
func GetUsers(c *fiber.Ctx) error {
    var users []models.Users
    database.DB.Db.Find(&users)
    return c.Status(fiber.StatusOK).JSON(users)
}

// CreateUser - Create a new user
func CreateUser(c *fiber.Ctx) error {
    user := new(models.Users)
    if err := c.BodyParser(user); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
    }
    database.DB.Db.Create(&user)
    return c.Status(fiber.StatusCreated).JSON(user)
}

// GetUser - Retrieve a user by ID
func GetUser(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
    }

    var user models.Users
    result := database.DB.Db.First(&user, id)
    if result.Error != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
    }
    return c.Status(fiber.StatusOK).JSON(user)
}

// UpdateUser - Update a user by ID
func UpdateUser(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
    }

    var user models.Users
    result := database.DB.Db.First(&user, id)
    if result.Error != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
    }

    if err := c.BodyParser(&user); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
    }

    database.DB.Db.Save(&user)
    return c.Status(fiber.StatusOK).JSON(user)
}

// DeleteUser - Delete a user by ID
func DeleteUser(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
    }

    result := database.DB.Db.Delete(&models.Users{}, id)
    if result.Error != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
    }

    return c.SendStatus(fiber.StatusNoContent)
}
