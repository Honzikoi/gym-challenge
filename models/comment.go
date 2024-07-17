package models

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Title        string    `gorm:"type:varchar(255)"`
	UserID       int       `gorm:"index"`
	CreationDate time.Time `gorm:"type:timestamp"`
	CreationTime time.Time `gorm:"type:timestamp"`
}
