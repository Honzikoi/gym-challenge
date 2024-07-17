package models

import (
	"gorm.io/gorm"
)

type Group struct {
    gorm.Model
    Name        string   `json:"name" gorm:"type:varchar(255)"`
    Description string   `json:"description" gorm:"type:text"`
    Users       []Users `gorm:"many2many:group_users;"` 
}
