package routes

import (
	"github.com/Honzikoi/gym-challenge/controllers"
	"github.com/gofiber/fiber/v2"
)

func SessionRoutes(app *fiber.App) {
	app.Get("/sessions", controllers.GetSessions)
	app.Get("/sessions/:id", controllers.GetSession)

	// app.Use(middlewares.RequireAuth)
	app.Post("/sessions", controllers.CreateSession)
	app.Put("/sessions/:id", controllers.UpdateSession)
	app.Delete("/sessions/:id", controllers.DeleteSession)
}
