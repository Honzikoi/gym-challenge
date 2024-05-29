package main

import (
	"github.com/Honzikoi/gym-challenge/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
	// app.Get("/singup", handlers.Singup)

}
