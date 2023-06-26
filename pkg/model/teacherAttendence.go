package model

import (
	"time"

	"gorm.io/gorm"
)

type TeacherAttendence struct {
	ID                uint      `gorm:"primary_key:autoIncrement"`
	TeacherAttendence string    `gorm:"not null;default:null"`
	TeacherID         uint      `gorm:"foreignKey"`
	CreatedAt         time.Time `gorm:"autoUpdate"`
	DeletedAt         gorm.DeletedAt
	UpdatedAt         time.Time `gorm:"autoUpdate"`
	Teacher           Teacher
	Date              time.Time `gorm:"type:time"`
}
