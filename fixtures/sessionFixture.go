package fixtures

import (
	"github.com/Honzikoi/gym-challenge/models"
	"gorm.io/gorm"
)

func LoadSessionFixtures(db *gorm.DB) error {
	sessions := []models.Sessions{
		// {Name: "Morning Workout", Description: "An energizing start to the day.", Group_id: 1},
		// {Name: "Evening Yoga", Description: "Relax and unwind with evening yoga.", Group_id: 2},
		// {Name: "Afternoon Cardio", Description: "Boost your heart rate with afternoon cardio.", Group_id: 3},
	}

	for _, session := range sessions {
		if err := db.Create(&session).Error; err != nil {
			return err
		}
	}

	return nil
}
