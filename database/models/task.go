package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title     string `gorm:"not null" json:"title"`
	Completed bool   `gorm:"default:false" json:"completed"`
	UserID    uint   `json:"user_id"`
	User      User   `gorm:"foreignKey:UserID" json:"-"`
}
