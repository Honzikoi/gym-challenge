package main

import (
	"github.com/Honzikoi/gym-challenge/handlers"
	"github.com/Honzikoi/gym-challenge/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
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
		app.Post("/signup", handlers.SignUp)
		app.Post("/login", handlers.Login)
		app.Post("/logout", handlers.Logout)
		app.Get("/swagger/*", swagger.HandlerDefault)

		// app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		// 	URL: "http://example.com/doc.json",
		// 	DeepLinking: false,
		// 	// Expand ("list") or Collapse ("none") tag groups by default
		// 	DocExpansion: "none",
		// 	// Prefill OAuth ClientId on Authorize popup
		// 	OAuth: &swagger.OAuthConfig{
		// 		AppName:  "OAuth Provider",
		// 		ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		// 	},
		// 	// Ability to change OAuth2 redirect uri location
		// 	OAuth2RedirectUrl: "http://localhost:3000/swagger/oauth2-redirect.html",
		// }))

		// Protected routes
		app.Use("/profile", middlewares.Protected())
		app.Get("/profile", handlers.Profile)
		// app.Get("/auth/google", handlers.GoogleLogin)
		// app.Get("/auth/google/callback", handlers.GoogleCallback)

	}
