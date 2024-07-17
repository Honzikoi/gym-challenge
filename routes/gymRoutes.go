package routes

import (
	"github.com/gofiber/fiber/v2"
)

func GymRoutes(app *fiber.App) {
	// Gym routes
	// app.Post("/gyms", middlewares.Protected(), middlewares.CheckRole("admin"), handlers.AddGym)
	// app.Post("/gyms/equipment", middlewares.Protected(), middlewares.CheckRole("coach"), handlers.AddEquipmentToGym)
}
