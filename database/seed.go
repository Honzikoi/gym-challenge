// database/seed.go
package database

import (
	"log"

	"github.com/Honzikoi/gym-challenge/models"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := seedUsers(tx); err != nil {
			return err
		}
		if err := seedGroups(tx); err != nil {
			return err
		}
		if err := seedSessions(tx); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Failed to seed data: %v", err)
	}
}

func seedUsers(db *gorm.DB) error {
	users := []models.Users{
		{Username: "test", Email: "test@test.test", Password: "test"},
	}
	return db.Create(&users).Error
}

func seedGroups(db *gorm.DB) error {
	groups := []models.Group{
		{Name: "Group 1"},
		{Name: "Group 2"},
		{Name: "Group 3"},
	}
	return db.Create(&groups).Error
}

func seedSessions(db *gorm.DB) error {
	sessions := []models.Sessions{
		{Model: gorm.Model{ID: 1}, Name: "Morning Session", Description: "Start your day with a burst of energy", Group_id: 1, Type: "Yoga"},
		{Model: gorm.Model{ID: 2}, Name: "Afternoon Session", Description: "Keep your momentum going", Group_id: 2, Type: "Muscu"},
		{Model: gorm.Model{ID: 3}, Name: "Evening Session", Description: "Wind down and relax", Group_id: 3, Type: "Cardio"},
	}
	return db.Create(&sessions).Error
}
