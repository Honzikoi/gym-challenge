package models

import (
	"gorm.io/gorm"
)

type Gym struct {
    gorm.Model
    Name      string      `gorm:"not null"`
    Location  string      `gorm:"not null"`
    Equipment []Equipment `gorm:"many2many:gym_equipments;"`
}

type Equipment struct {
    gorm.Model
    Name string `gorm:"not null"`
}