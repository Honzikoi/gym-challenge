package routes

import (
	"github.com/Honzikoi/gym-challenge/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Other routes...

	app.Post("/workouts", controllers.CreateWorkout)
	app.Get("/workouts", controllers.GetWorkouts)
	app.Get("/workouts/:id", controllers.GetWorkout)
	app.Put("/workouts/:id", controllers.UpdateWorkout)
	app.Delete("/workouts/:id", controllers.DeleteWorkout)
}
