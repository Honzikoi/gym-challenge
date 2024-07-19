package main

import (
	"github.com/Honzikoi/gym-challenge/api/handlers"
	"github.com/Honzikoi/gym-challenge/api/middlewares"
	"github.com/Honzikoi/gym-challenge/routes"
	"github.com/gofiber/fiber/v2"
)

//	@title			Gym API
//	@version		1.0
//	@description	Gym API for users.
//	@termsOfService	http://swagger.io/terms/
//	@contact.name	API Support
//	@contact.email	fiber@swagger.io
//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//	@host			localhost:3000
//	@BasePath		/

// SetupRoutes sets up all the routes for the application
func setupRoutes(app *fiber.App) {
	// Public routes
	app.Get("/", handlers.Home)

	routes.ConnectionRoutes(app)
	routes.GymRoutes(app)
	routes.GroupRoutes(app)
	routes.UserRoutes(app)
	routes.WorkoutRoutes(app)
	routes.SwaggerRoutes(app)
	routes.ConnectionRoutes(app)
	routes.AboutRoutes(app)
	routes.SessionRoutes(app)
	routes.CoachRoutes(app)
	routes.RoleRoutes(app)

	// Protected routes
	app.Use("/profile", middlewares.Protected())
	app.Get("/profile", handlers.Profile)
	// app.Get("/auth/google", handlers.GoogleLogin)
	// app.Get("/auth/google/callback", handlers.GoogleCallback)

}
