package model

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID        uint `gorm:"primary_key:autoIncrement"`
	Role      string
	Status    string
	User      []*Login
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
	UpdatedAt time.Time
}
