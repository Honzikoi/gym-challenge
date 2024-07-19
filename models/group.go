package models

import (
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	SessionID uint       `json:"session_id"`
	Name      string     `json:"name" gorm:"type:varchar(255)"`
	Users     []Users    `json:"users" gorm:"many2many:user_groups;"`
	Sessions  []Sessions `json:"sessions"`
}
