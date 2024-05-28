package main

import (
	"github.com/Honzikoi/gym-challenge/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
	// app.Get("/facts", handlers.GetFacts)
	// app.Get("/facts/:id", handlers.GetFact)
	// app.Post("/facts", handlers.NewFact)
	// app.Delete("/facts/:id", handlers.DeleteFact)
	// app.Put("/facts/:id", handlers.UpdateFact
}
