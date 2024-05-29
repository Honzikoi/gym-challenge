package main

import (
	"github.com/Honzikoi/gym-challenge/config"
	"github.com/Honzikoi/gym-challenge/database"
	"github.com/Honzikoi/gym-challenge/fixtures"
	"github.com/Honzikoi/gym-challenge/handlers"
	"github.com/Honzikoi/gym-challenge/middlewares"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()
	fixtures.CreateUserFixture()
	jwt := middlewares.NewAuthMiddleware(config.Secret)
	// Create a Login route
	app.Post("/login", handlers.Login)
	// Create a protected route
	app.Get("/protected", jwt, handlers.Protected)
	setupRoutes(app)

	app.Listen(":3000")
}
