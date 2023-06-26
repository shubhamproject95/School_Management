package model

import (
	"time"
)

type ClassSubject struct {
	ID        uint `gorm:"primary_key:autoIncrement"`
	ClassID   uint `gorm:"foreignKey"`
	SubjectID uint `gorm:"foreignKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
