package models

import (
	"gorm.io/gorm"
)

type LoginRequest struct {
	gorm.Model
	Email    string `json:"email" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
}

type LoginResponse struct {
	gorm.Model
	Token string `json:"token" gorm:"not null"`
}
