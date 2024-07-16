package models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Email    string `json:"email" gorm:"text;not null"`
	Username string `json:"username" gorm:"text;not null"`
	Password string `json:"password" gorm:"text;not null"`
	Roles  []Role   `json:"role_id" gorm:"foreignKey:UserID;references:ID"`
}

type Role struct {
	gorm.Model
	Name        string `json:"name" gorm:"text;not null"`
	Description string `json:"description" gorm:"text;not null"`
	Permissions string `json:"permissions" gorm:"text;not null"`
	Users []Users `gorm:"many2many:user_roles;"`
}
