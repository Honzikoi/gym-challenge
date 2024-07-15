package models

import (
	"gorm.io/gorm"
)

type Group struct {
    gorm.Model
    Name        string
    Description string
    Coaches     []Users `gorm:"many2many:group_coaches;"`
    Members     []Users `gorm:"many2many:user_groups;"`
}

type UserGroup struct {
    UserID  uint
    GroupID uint
    Status  string // "pending", "approved", "rejected"
}
