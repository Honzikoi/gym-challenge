package controllers

import (
	"github.com/Honzikoi/gym-challenge/database"
	"github.com/Honzikoi/gym-challenge/models"
	"github.com/gofiber/fiber/v2"
)

func CreateWorkout(c *fiber.Ctx) error {
	workout := new(models.Workout)
	if err := c.BodyParser(workout); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	if err := database.DB.Db.Create(&workout).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot create workout",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(workout)
}

func GetWorkouts(c *fiber.Ctx) error {
	var workouts []models.Workout
	if err := database.DB.Db.Find(&workouts).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot get workouts",
		})
	}
	return c.JSON(workouts)
}

func GetWorkout(c *fiber.Ctx) error {
	id := c.Params("id")
	var workout models.Workout
	if err := database.DB.Db.First(&workout, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Workout not found",
		})
	}
	return c.JSON(workout)
}

func UpdateWorkout(c *fiber.Ctx) error {
	id := c.Params("id")
	var workout models.Workout
	if err := database.DB.Db.First(&workout, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Workout not found",
		})
	}

	if err := c.BodyParser(&workout); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	if err := database.DB.Db.Save(&workout).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot update workout",
		})
	}
	return c.JSON(workout)
}

func DeleteWorkout(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := database.DB.Db.Delete(&models.Workout{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot delete workout",
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
