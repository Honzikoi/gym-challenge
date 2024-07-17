package models

type UserGroup struct {
	UserID  uint `gorm:"primaryKey"`
	GroupID uint `gorm:"primaryKey"`
}
