package routes

import (
	"github.com/Honzikoi/gym-challenge/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func ConnectionRoutes(app *fiber.App) {

	app.Post("/signup", handlers.SignUp)
	app.Post("/login", handlers.Login)
	app.Post("/logout", handlers.Logout)
}