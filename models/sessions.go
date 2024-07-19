package models

import (
	"time"

	"gorm.io/gorm"
)

type Sessions struct {
	gorm.Model
	SessionID         uint      `json:"session_id"`
	WorkOutTitle      string    `json:"workOutTitle" gorm:"not null"`
	TimeLeftInHour    int       `json:"timeLeftInHour"`
	MovesNumber       int       `json:"movesNumber"`
	SetsNumber        int       `json:"setsNumber"`
	DurationInMinutes int       `json:"durationInMinutes"`
	Description       string    `json:"description"`
	Date              time.Time `json:"date"`
	Localisation      string    `json:"localisation"`
	ExerciseType      string    `json:"exerciseType"`
	IsActive          bool      `json:"isactive"`
	Reviews           string    `json:"reviews"`
	Rating            int       `json:"rating"`
	IsWorkoutOfDay    bool      `json:"isWorkoutOfDay"`
	Name              string    `json:"name"`
	GroupID           uint      `json:"group_id"`
	Group             Group     `json:"group" gorm:"foreignKey:GroupID"`
}
