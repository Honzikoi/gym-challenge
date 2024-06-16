package main

import (
	"github.com/Honzikoi/gym-challenge/database"
	"github.com/Honzikoi/gym-challenge/fixtures"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()
	fixtures.CreateUserFixture()


	setupRoutes(app)

	app.Listen(":3000")
}
