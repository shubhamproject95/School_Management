package model

import (
	"time"
)

type ClassTeacher struct {
	ID        uint `gorm:"primary_key"`
	ClassID   uint `gorm:"foreignKey"`
	TeacherID uint `gorm:"foreignKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
