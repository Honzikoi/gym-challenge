package routes

import (
	"github.com/Honzikoi/gym-challenge/controllers"
	"github.com/gofiber/fiber/v2"
)

func CoachRoutes(app *fiber.App) {
	app.Get("/coaches", controllers.GetUsers)
	app.Get("/coaches/:id", controllers.GetUser)
	app.Post("/coaches", controllers.CreateCoach)
	app.Put("/coaches/:id", controllers.UpdateUser)
	app.Delete("/coaches/:id", controllers.DeleteUser)
}
