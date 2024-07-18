package routes

import (
	"github.com/Honzikoi/gym-challenge/controllers"
	"github.com/gofiber/fiber/v2"
)

func GymRoutes(app *fiber.App) {
	app.Get("/gyms", controllers.GetGyms)
	app.Get("/gyms/:id", controllers.GetGym)
	app.Post("/gyms", controllers.CreateGym)
	app.Put("/gyms/:id", controllers.UpdateGym)
	app.Delete("/gyms/:id", controllers.DeleteGym)
}
