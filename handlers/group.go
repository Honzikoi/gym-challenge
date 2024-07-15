package handlers

import (
	"github.com/Honzikoi/gym-challenge/database"
	"github.com/Honzikoi/gym-challenge/models"
	"github.com/gofiber/fiber/v2"
)

func RequestJoinGroup(c *fiber.Ctx) error {
    type Request struct {
        GroupID uint `json:"group_id"`
        UserID  uint `json:"user_id"`
    }

    var request Request
    if err := c.BodyParser(&request); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
    }

    userGroup := models.UserGroup{
        UserID:  request.UserID,
        GroupID: request.GroupID,
        Status:  "pending",
    }

    database.DB.Db.Create(&userGroup)
    return c.JSON(userGroup)
}

func ManageJoinRequests(c *fiber.Ctx) error {
    type Request struct {
        UserGroupID uint   `json:"user_group_id"`
        Action      string `json:"action"` // "approve" or "reject"
    }

    var request Request
    if err := c.BodyParser(&request); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
    }

    var userGroup models.UserGroup
    if err := database.DB.Db.First(&userGroup, request.UserGroupID).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Request not found"})
    }

    if request.Action == "approve" {
        userGroup.Status = "approved"
    } else if request.Action == "reject" {
        database.DB.Db.Delete(&userGroup)
        return c.JSON(fiber.Map{"message": "Request rejected"})
    } else {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid action"})
    }

    database.DB.Db.Save(&userGroup)
    return c.JSON(userGroup)
}
