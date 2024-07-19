package models

import (
	"gorm.io/gorm"
)

type Workout struct {
	gorm.Model
	WorkOutTitle      string  `json:"workOutTitle"`
	ImagePath         string  `json:"imagePath"`
	TimeLeftInHour    int     `json:"timeLeftInHour"`
	MovesNumber       int     `json:"movesNumber"`
	SetsNumber        int     `json:"setsNumber"`
	DurationInMinutes int     `json:"durationInMinutes"`
	Rating            int     `json:"rating"`
	Description       string  `json:"description"`
	Reviews           string  `json:"reviews"`
	Comments          string  `json:"comments"`
	PriceInDollars    float64 `json:"priceInDollars"`
	HasFreeTrial      bool    `json:"hasFreeTrial"`
	IsWorkoutOfDay    bool    `json:"isWorkoutOfDay"`
	SessionID         uint    `json:"sessionID"`
}
