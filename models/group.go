// models/groups.go
package models

import (
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	Name        string  `json:"name" gorm:"type:varchar(255)"`
	Description string  `json:"description" gorm:"type:text"`
	Users       []Users `json:"users" gorm:"many2many:user_groups;"`
	Status      string  `json:"status" gorm:"type:varchar(255)"`
}
