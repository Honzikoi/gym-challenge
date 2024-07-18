package models

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	RoleID      uint    `json:"role_id"`
	Name        string  `json:"name" gorm:"text;not null"`
	Description string  `json:"description" gorm:"text;not null"`
	Permissions string  `json:"permissions" gorm:"text;not null"`
	Users       []Users `json:"users" gorm:"foreignKey:UserID"`
}
