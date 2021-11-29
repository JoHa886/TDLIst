package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	User      User   `gorm:"foreignKey:Uid"`
	Uid       uint   `gorm:"not null"`
	Title     string `gorm:"index,not null"`
	Status    int    `gorm:"default:0"`
	Content   string `gorm:"type:longtext"`
	StartTime int64
	EndTime   int64
}
