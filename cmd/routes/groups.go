package routes

import (
	"github.com/Honzikoi/gym-challenge/handlers"
	"github.com/Honzikoi/gym-challenge/middlewares"
	"github.com/gofiber/fiber/v2"
)

func GroupRoutes(app *fiber.App) {
	// Group routes
	// app.Post("/groups", middlewares.Protected(), middlewares.CheckRole("coach"), handlers.CreateGroup)
	app.Post("/groups/join", middlewares.Protected(), middlewares.CheckRole("user"), handlers.RequestJoinGroup)
	app.Post("/groups/manage", middlewares.Protected(), middlewares.CheckRole("coach"), handlers.ManageJoinRequests)
	app.Post("/roles/assign", middlewares.Protected(), middlewares.CheckRole("admin"), handlers.AssignRole)
}
