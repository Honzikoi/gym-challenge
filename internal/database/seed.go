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
		// if err := seedGroups(tx); err != nil {
		// 	return err
		// }
		if err := seedUserGroups(tx); err != nil {
			return err
		}
		// if err := seedSessions(tx); err != nil {
		// 	return err
		// }
		return nil
	})

	if err != nil {
		log.Fatalf("Failed to seed data: %v", err)
	}
}

func seedUsers(db *gorm.DB) error {
	users := []models.Users{
		{Username: "test1", Email: "test1@test.test", Password: "test"},
		{Username: "test2", Email: "test2@test.test", Password: "test"},
		{Username: "test3", Email: "test3@test.test", Password: "test"},
	}
	return db.Create(&users).Error
}

// func seedGroups(db *gorm.DB) error {
// 	groups := []models.Group{
// 		{Name: "Group 1", Description: "Description for Group 1", Status: "Active"},
// 		{Name: "Group 2", Description: "Description for Group 2", Status: "Inactive"},
// 		{Name: "Group 3", Description: "Description for Group 3", Status: "Pending"},
// 	}
// 	return db.Create(&groups).Error
// }

func seedUserGroups(db *gorm.DB) error {
	userGroups := []models.UserGroup{
		{UserID: 1, GroupID: 1},
		{UserID: 2, GroupID: 1},
		{UserID: 3, GroupID: 1},
		{UserID: 1, GroupID: 2},
		{UserID: 2, GroupID: 2},
		{UserID: 1, GroupID: 3},
	}
	return db.Create(&userGroups).Error
}

// func seedSessions(db *gorm.DB) error {
// 	sessions := []models.Sessions{
// 		{Model: gorm.Model{ID: 1}, Name: "Morning Session", Description: "Start your day with a burst of energy", Group_id: 1, Type: "Yoga"},
// 		{Model: gorm.Model{ID: 2}, Name: "Afternoon Session", Description: "Keep your momentum going", Group_id: 2, Type: "Muscu"},
// 		{Model: gorm.Model{ID: 3}, Name: "Evening Session", Description: "Wind down and relax", Group_id: 3, Type: "Cardio"},
// 	}
// 	return db.Create(&sessions).Error
// }
