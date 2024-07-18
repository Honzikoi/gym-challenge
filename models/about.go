package models

import (
	"time"

	"gorm.io/gorm"
)

type About struct {
	gorm.Model
	AboutID     uint      `json:"about_id"`
	Name        string    `json:"name" gorm:"type:varchar(255)"`
	Surname     string    `json:"surname" gorm:"type:varchar(255)"`
	Description string    `json:"description" gorm:"type:text"`
	Notes       string    `json:"notes" gorm:"type:text"`
	Address     string    `json:"address" gorm:"type:varchar(255)"`
	Gender      string    `json:"gender" gorm:"type:varchar(255)"`
	Age         time.Time `json:"age" gorm:"type:timestamp"`
}
