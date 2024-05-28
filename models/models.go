package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Email    string `json:"email" gorm:"text;not null"`
	Username string `json:"username" gorm:"text;not null"`
	Password string `json:"password" gorm:"text;not null"`
	Role_id  int    `json:"role_id" gorm:"int;not null;default:1"`
}

type Groups struct {
	gorm.Model
	Users_id    int    `json:"users_id" gorm:"int;not null;default:"`
	Name        string `json:"name" gorm:"text;not null"`
	Coach_id    int    `json:"coach_id" gorm:"int"`
	Description string `json:"description" gorm:"text:not null"`
}

type Sessions struct {
	gorm.Model
	Name        string `json:"name" gorm:"text;not null;default:null"`
	Description string `json:"user_id" gorm:"text;not null"`
	Group_id    int    `json:"group_id" gorm:"int;not null"`
}
