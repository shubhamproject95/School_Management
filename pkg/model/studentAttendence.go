package model

import (
	"time"

	"gorm.io/gorm"
)

type StudentAttendence struct {
	gorm.Model
	ID         uint      `gorm:"primary_key:autoIncrement"`
	Attendence string    `gorm:"not null;default:null"`
	StudentID  uint      `gorm:"foreignKey"`
	Date       time.Time `gorm:"type:time"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
