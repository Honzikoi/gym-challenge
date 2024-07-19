package models

import (
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	SessionID uint    `json:"session_id"`
	Name      string  `json:"name" gorm:"type:varchar(255)"`
	UserID    uint    `json:"user_id"`
	User      Users   `json:"user_creator" gorm:"foreignKey:UserID"`
	Users     []Users `json:"users" gorm:"many2many:user_groups;"`
}
