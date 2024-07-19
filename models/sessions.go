package models

import (
	"time"

	"gorm.io/gorm"
)

type Sessions struct {
	gorm.Model
	WorkOutTitle      string    `json:"workOutTitle" gorm:"not null"`
	TimeLeftInHour    string    `json:"timeLeftInHour"`
	MovesNumber       string    `json:"movesNumber"`
	SetsNumber        string    `json:"setsNumber"`
	DurationInMinutes string    `json:"durationInMinutes"`
	Description       string    `json:"description"`
	Date              time.Time `json:"date"`
	Localisation      string    `json:"localisation"`
	ExerciseType      string    `json:"exerciseType"`
	IsActive          bool      `json:"isactive"`
	Reviews           string    `json:"reviews"`
	Rating            string    `json:"rating"`
	IsWorkoutOfDay    bool      `json:"isWorkoutOfDay"`
}
