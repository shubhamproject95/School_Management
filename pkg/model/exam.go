package model

import (
	"time"

	"gorm.io/gorm"
)

type Exam struct {
	ID             uint `gorm:"primary_key:autoIncrement"`
	SubjectID      uint `gorm:"foreignKey"`
	ClassID        uint `gorm:"foreignKey"`
	ExamType       string
	InternalMarks  uint
	ExtenalMarks   uint
	PracticalMarks uint
	StudentExam    []StudentExam
	CreatedAt      time.Time `gorm:"autoUpdate"`
	UpdatedAt      time.Time `gorm:"autoUpdate"`
	DeletedAt      gorm.DeletedAt
}
