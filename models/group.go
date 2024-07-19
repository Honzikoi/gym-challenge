// models/groups.go
package models

import (
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	GroupID uint    `json:"group_id"`
	Name    string  `json:"name" gorm:"type:varchar(255)"`
	Users   []Users `json:"users" gorm:"many2many:user_groups;"`
}
