package model

import "time"

type Tour struct {
	ID          uint `gorm:"primary_key;autoIncrement"`
	StudentID   uint `gorm:"foreignKey"`
	ClassID     uint `gorm:"foreignKey"`
	TeacherID   uint `gorm:"foreignKey"`
	Destination string
	Date        string
	Cost        uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
