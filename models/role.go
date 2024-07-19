package models

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name        string `json:"name" gorm:"type:text;not null"`
	Description string `json:"description" gorm:"type:text;not null"`
	Permissions string `json:"permissions" gorm:"type:text;not null"`
	Users       []Users `json:"users" gorm:"foreignKey:RoleID"`
}
