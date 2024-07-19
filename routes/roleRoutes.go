package routes

import (
	"github.com/Honzikoi/gym-challenge/controllers"
	"github.com/gofiber/fiber/v2"
)

func RoleRoutes(app *fiber.App) {
	app.Get("/roles", controllers.GetRoles)
	app.Get("/roles/:id", controllers.GetRole)
	app.Post("/roles", controllers.CreateRole)
	app.Put("/roles/:id", controllers.UpdateRole)
	app.Delete("/roles/:id", controllers.DeleteRole)
}
