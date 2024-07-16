package models

import (
	"gorm.io/gorm"
)

type Sessions struct {
	gorm.Model
	Name        string    `json:"name" gorm:"text;not null;default:null"`
	Description string    `json:"user_id" gorm:"text;not null"`
	Group_id    int       `json:"group_id" gorm:"int;not null"`
	Workouts    []Workout `json:"workouts"`
}

type UserGroup struct {
	UserID  uint
	GroupID uint
	Status  string // "pending", "approved", "rejected"
}
