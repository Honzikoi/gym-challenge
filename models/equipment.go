package models

import (
	"gorm.io/gorm"
)

type Equipment struct {
	gorm.Model
    Name    string `gorm:"type:varchar(255)"`
    Purpose string `gorm:"type:text"`
}
