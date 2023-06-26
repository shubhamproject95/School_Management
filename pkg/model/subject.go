package model

import (
	"time"

	"gorm.io/gorm"
)

type Subject struct {
	gorm.Model
	ID        uint      `gorm:"primary_key:autoIncrement"`
	Subject   string    `gorm:"not null;default:null"`
	CreatedAt time.Time `gorm:"autoUpdate"`
	UpdatedAt time.Time `gorm:"autoUpdate"`

	Exam         Exam
	ClassSubject *ClassSubject
	Homework     *Homework
}
