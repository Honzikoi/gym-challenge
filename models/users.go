package models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Email    string `json:"email" gorm:"text;not null"`
	Username string `json:"username" gorm:"text;not null"`
	Password string `json:"password" gorm:"text;not null"`
	Role_id  int    `json:"role_id" gorm:"int;not null;default:1"`
}
