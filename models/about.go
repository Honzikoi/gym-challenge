package models

import (
	"gorm.io/gorm"
)

type About struct {
	gorm.Model
	Name    string `json:"name" gorm:"text"`
	Surname string `json:"surname" gorm:"text"`
}
