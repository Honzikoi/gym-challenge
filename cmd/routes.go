package main

import (
	"os"

	"github.com/Honzikoi/gym-challenge/handlers"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes sets up all the routes for the application
func setupRoutes(app *fiber.App) {
	// Public routes
	app.Get("/", handlers.Home)
	app.Post("/signup", handlers.SignUp)
	app.Post("/login", handlers.Login)
	app.Get("/auth/google", handlers.GoogleLogin)
	app.Get("/auth/google/callback", handlers.GoogleCallback)

	// JWT Protected routes
	protected := app.Group("/", jwt.New(jwt.Config{
		SigningKey: []byte(os.Getenv("SECRET_KEY")),
	}))

	protected.Get("/profile", handlers.Profile)
}