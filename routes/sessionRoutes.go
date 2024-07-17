package routes

import (
	"github.com/Honzikoi/gym-challenge/controllers"
	"github.com/Honzikoi/gym-challenge/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SessionRoutes(app *fiber.App) {
    app.Get("/sessions", controllers.GetSessions)
	app.Get("/sessions/:id", controllers.GetSession)

	app.Post("/sessions", middlewares.IsAdminOrCoach, controllers.CreateSession)
    app.Put("/sessions/:id", middlewares.IsAdminOrCoach, controllers.UpdateSession)
    app.Delete("/sessions/:id", middlewares.IsAdminOrCoach, controllers.DeleteSession)
}
