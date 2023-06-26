package model

import "time"

type StudentHomework struct {
	ID         uint `gorm:"primary_key:autoIncrement"`
	HomeworkID uint `gorm:"foreignKey"`
	StudentID  uint `gorm:"foreignKey"`
	CreatedAt  time.Time
	DeletedAt  time.Time
	UpdatedAt  time.Time
}
