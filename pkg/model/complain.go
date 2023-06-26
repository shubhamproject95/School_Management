package model

import (
	"time"

	"gorm.io/gorm"
)

type Complain struct {
	ID        uint   `gorm:"primary_key:autoIncrement"`
	Complain  string `gorm:"not null;default:null"`
	StudentID uint   `gorm:"foreignKey"`
	Student   Student
	CreatedAt time.Time `gorm:"autoUpdate"`
	UpdatedAt time.Time `gorm:"autoUpdate"`
	DeletedAt gorm.DeletedAt
}
