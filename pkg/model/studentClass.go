package model

import "gorm.io/gorm"

type StudentClass struct {
	gorm.Model
	ID            uint `gorm:"primary_key:autoIncrement"`
	StudentID     uint `gorm:"foreignKey"`
	ClassID       uint `gorm:"foreignKey"`
	SessionYearID uint `gorm:"foreignKey"`
}
