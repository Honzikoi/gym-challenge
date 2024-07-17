package main

import (
	"log"
	"os"

	"github.com/Honzikoi/gym-challenge/database"
	"github.com/gofiber/fiber/v2"
)

// @title           Swagger GoCoach-API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3000
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

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
