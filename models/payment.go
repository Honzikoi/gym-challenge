package models

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	CardName   string `gorm:"type:varchar(255)"`
	CardNumber int
	CardPin    int
	ExpDate    time.Time `gorm:"type:timestamp"`
}
