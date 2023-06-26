package model

import (
	"time"

	"gorm.io/gorm"
)

type Homework struct {
	ID             uint `gorm:"primary_key:autoIncrement"`
	SubjectID      uint `gorm:"foreignKey"`
	TeacherID      uint `gorm:"foreignKey"`
	ClassID        uint `gorm:"foreignKey"`
	Description    string
	SubmissionDate string
	CreatedAt      time.Time `gorm:"autoUpdate"`
	UpdatedAt      time.Time `gorm:"autoUpdate"`
	gorm.DeletedAt
	StudentHomework *StudentHomework
}
