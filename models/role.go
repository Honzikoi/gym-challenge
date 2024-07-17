package models

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name        string  `json:"name" gorm:"text;not null"`
	Description string  `json:"description" gorm:"text;not null"`
	Permissions string  `json:"permissions" gorm:"text;not null"`
	Users       []Users `gorm:"many2many:user_roles;"`
}
