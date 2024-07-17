package fixtures

import (
	"github.com/Honzikoi/gym-challenge/models"
	"gorm.io/gorm"
)

func LoadWorkoutFixtures(db *gorm.DB) error {
	workouts := []models.Workout{
		{WorkOutTitle: "Full Body Blast", ImagePath: "/images/full_body_blast.png", TimeLeftInHour: 1, MovesNumber: 10, SetsNumber: 3, DurationInMinutes: 60, Rating: 5, Description: "A full-body workout to build strength and endurance.", Reviews: "Amazing workout!", Comments: "Loved it!", PriceInDollars: 29.99, HasFreeTrial: true, IsWorkoutOfDay: true, SessionID: 1},
		{WorkOutTitle: "Cardio Rush", ImagePath: "/images/cardio_rush.png", TimeLeftInHour: 2, MovesNumber: 8, SetsNumber: 4, DurationInMinutes: 45, Rating: 4, Description: "High-intensity cardio to get your heart pumping.", Reviews: "Great for burning calories!", Comments: "A bit tough but worth it!", PriceInDollars: 19.99, HasFreeTrial: false, IsWorkoutOfDay: false, SessionID: 3},
		{WorkOutTitle: "Yoga Flow", ImagePath: "/images/yoga_flow.png", TimeLeftInHour: 3, MovesNumber: 5, SetsNumber: 2, DurationInMinutes: 30, Rating: 5, Description: "Calm your mind and body with a gentle yoga flow.", Reviews: "So relaxing!", Comments: "Perfect for beginners.", PriceInDollars: 14.99, HasFreeTrial: true, IsWorkoutOfDay: false, SessionID: 2},
	}

	for _, workout := range workouts {
		if err := db.Create(&workout).Error; err != nil {
			return err
		}
	}

	return nil
}
