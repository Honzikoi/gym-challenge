package models

import (
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	SessionID uint   `json:"session_id"`
	Name      string `json:"name" gorm:"type:varchar(255)"`
	UserID    uint   `json:"user_id"`
}
