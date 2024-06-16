package main

import (
	"github.com/Honzikoi/gym-challenge/handlers"
	"github.com/Honzikoi/gym-challenge/middlewares"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes sets up all the routes for the application
func setupRoutes(app *fiber.App) {
	// Public routes
	app.Get("/", handlers.Home)
	app.Post("/signup", handlers.SignUp)
	app.Post("/login", handlers.Login)
	// Add your middleware to protect the profile route
	app.Use("/profile", middlewares.Protected())

	// Register the profile route
	app.Get("/profile", handlers.Profile)
	// app.Get("/auth/google", handlers.GoogleLogin)
	// app.Get("/auth/google/callback", handlers.GoogleCallback)

}
