package main

import (
	"log"
	"os"

	_ "github.com/Honzikoi/gym-challenge/cmd/docs"
	"github.com/Honzikoi/gym-challenge/database"
	"github.com/gofiber/fiber/v2"
)

// @title           Swagger GoCoach-API
// @version         1.0
// @description     This is a sample server celler server.

// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3000
// @BasePath  /

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

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
