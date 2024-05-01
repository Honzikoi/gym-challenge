package main

import (
	"github.com/Honzikoi/gym-challenge/database"
	"github.com/gofiber/fiber/v3"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoute(app)
	
	app.Listen(":3000")
}