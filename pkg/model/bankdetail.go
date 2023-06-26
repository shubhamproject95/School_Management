package model

import (
	"time"

	"gorm.io/gorm"
)

type BankDetail struct {
	ID          uint   `gorm:"autoIncrement;PRIMARY_KEY"`
	StaffID     uint   `gorm:"foreignKey"`
	Name        string `gorm:"not null;default:null"`
	Bank        string `gorm:"not null;default:null"`
	BankAccount string `gorm:"not null;default:null"`
	IFSC        string `gorm:"not null;default:null"`
	BranchCode  string `gorm:"not null;default:null"`
	IsDefault   bool   `gorm:"default:true"`
	Salary      Salary
	CreatedAt   time.Time
	DeletedAt   gorm.DeletedAt
	UpdatedAt   time.Time
}
