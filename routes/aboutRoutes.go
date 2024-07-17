package routes

import (
	"github.com/Honzikoi/gym-challenge/controllers"
	"github.com/gofiber/fiber/v2"
)

func AboutRoutes(app *fiber.App) {
    app.Get("/abouts", controllers.GetAbouts)
    app.Post("/abouts", controllers.CreateAbout)
    app.Get("/abouts/:id", controllers.GetAbout)
    app.Put("/abouts/:id", controllers.UpdateAbout)
    app.Delete("/abouts/:id", controllers.DeleteAbout)
}
