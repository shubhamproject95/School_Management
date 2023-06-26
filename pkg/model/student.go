package model

import (
	"time"
)

type Student struct {
	ID                uint   `gorm:"primary_key:autoIncrement:"`
	Name              string `gorm:"not null;default:null"`
	Address           string `gorm:"not null;default:null"`
	DOB               string `gorm:"not null;default:null"`
	Father_Name       string `gorm:"not null;default:null"`
	Mother_Name       string `gorm:"not null;default:null"`
	Status            string `gorm:"not null;default:null"`
	StudentAttendence []StudentAttendence
	Guardian          []Guardian
	StudentClass      StudentClass
	Complain          []Complain
	StudentExam       []StudentExam
	Fees              []Fees
	StudentHomework   StudentHomework
	Tour              []Tour

	CreatedAt time.Time `gorm:"autoUpdate"`
	UpdatedAt time.Time `gorm:"autoUpdate"`
}
