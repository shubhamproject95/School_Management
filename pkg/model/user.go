package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID             uint `gorm:"primary_key:autoIncrement"`
	UserName       string
	Password       string
	UserType       string
	RolesID        uint `gorm:"foreignKey"`
	ResetPassword  ResetPassword
	UpdatePassword UpdatePassword
	CreatedAt      time.Time
	DeletedAt      gorm.DeletedAt
	UpdatedAt      time.Time
}
