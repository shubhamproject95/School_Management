package model

import (
	"time"

	"gorm.io/gorm"
)

type Result struct {
	ID         uint `gorm:"primary_key:autoIncrement"`
	ResultType string
	RollNumber int
	//Subject    Subject
	// Student Student
	// Class   Class
	// Teacher Teacher
	CreatedAt time.Time
	UpdatedAt time.Time
	DelatedAt gorm.DeletedAt
}
