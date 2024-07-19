package models

import "gorm.io/gorm"

type UserGroup struct {
	gorm.Model
	UserID  uint `gorm:"primaryKey"`
	GroupID uint `gorm:"primaryKey"`
}
