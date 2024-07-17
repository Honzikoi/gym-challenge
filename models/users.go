package models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Email    string `json:"email" gorm:"type:text;not null"`
	Username string `json:"username" gorm:"type:text;not null"`
	Password string `json:"password" gorm:"type:text;not null"`
	About    About  `json:"about" gorm:"foreignKey:UserID"`
}
