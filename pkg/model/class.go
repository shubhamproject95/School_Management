package model

import (
	"time"

	"gorm.io/gorm"
)

type Class struct {
	ID            uint      `gorm:"primary_key:autoIncrement"`
	Class         string    `gorm:"unique:not null;default:null"`
	CreatedAt     time.Time `gorm:"autoUpdate"`
	UpdatedAt     time.Time `gorm:"autoUpdate"`
	DeletedAt     gorm.DeletedAt
	StudentExam   *StudentExam
	Exam          []*Exam
	ClassHomework []*Homework
	ClassSubject  []*ClassSubject
	ClassTeacher  []*ClassTeacher
	StudentClass  []*StudentClass
	Fees          []*Fees
	Tour          []Tour
}
