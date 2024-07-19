package models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	UserID   uint    `json:"user_id"`
	Email    string  `json:"email" gorm:"type:text;not null"`
	Username string  `json:"username" gorm:"type:text;not null"`
	Password string  `json:"password" gorm:"type:text;not null"`
	AboutID  uint    `json:"about_id"`
	About    About   `json:"about" gorm:"foreignKey:AboutID"`
	RoleID   uint    `json:"role_id" gorm:"default:1"`
	Role     Role    `json:"role" gorm:"foreignKey:RoleID"`
	Groups   []Group `json:"groups" gorm:"many2many:user_groups;"`
}
