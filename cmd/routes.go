package main

import (
	"github.com/Honzikoi/gym-challenge/handlers"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes sets up all the routes for the application
func setupRoutes(app *fiber.App) {
	// Public routes
	app.Get("/", handlers.Home)
	app.Post("/signup", handlers.SignUp)
	app.Post("/login", handlers.Login)
	app.Get("/profile", handlers.Profile)
	// app.Get("/auth/google", handlers.GoogleLogin)
	// app.Get("/auth/google/callback", handlers.GoogleCallback)

}
