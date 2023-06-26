package model

import (
	"time"

	"gorm.io/gorm"
)

type StaffType string

const (
	StaffTeacher StaffType = "Teacher"
	StaffPeon    StaffType = "Peon"
)

type Staff struct {
	ID          uint   `gorm:"autoIncrement;faker:customIdFaker"`
	Name        string `gorm:"not null;default:null"`
	DOB         string `gorm:"not null;default:null"`
	JoiningDate string `gorm:"not null;default:null"`
	Aadharno    string `gorm:"not null;default:null"`
	StaffType   string `gorm:"not null;default:null"`
	Salary      Salary
	BankDetail  BankDetail
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
