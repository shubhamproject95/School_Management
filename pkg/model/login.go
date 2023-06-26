package model

import (
	"time"

	"gorm.io/gorm"
)

type UserType string

const (
	AdminUser   UserType = "Admin"
	TeacherUser UserType = "Teacher"
	StudentUser UserType = "Student"
)

type Login struct {
	ID         uint   `gorm:"primary_key:autoIncrement"`
	UserName   string `gorm:"unique"`
	Password   string
	UserType   string
	UserTypeID uint
	RoleID     uint
	Token      string
	Role       *Role //`gorm:"foreignKey:RoleID"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}
