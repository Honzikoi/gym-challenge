package main

import (
	"log"
	"os"

	_ "github.com/Honzikoi/gym-challenge/docs"
	"github.com/Honzikoi/gym-challenge/internal/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	// Seed the database
	// database.Seed(database.DB.Db) Temp Comment

	// Get the port from environment variable or default to 3000
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Start the server
	log.Fatal(app.Listen(":" + port))
}
