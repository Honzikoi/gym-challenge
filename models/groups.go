package models

import (
	"gorm.io/gorm"
)

type Groups struct {
	gorm.Model
	Users_id    int    `json:"users_id" gorm:"int;not null;default:"`
	Name        string `json:"name" gorm:"text;not null"`
	Coach_id    int    `json:"coach_id" gorm:"int"`
	Description string `json:"description" gorm:"text:not null"`
}
