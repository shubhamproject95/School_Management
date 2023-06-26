package model

import (
	"time"

	"gorm.io/gorm"
)

type GuardianType string

const (
	GuardianFather   GuardianType = "Father"
	GuardianMother   GuardianType = "Mother"
	GuardianSibling  GuardianType = "Sibling"
	GuardianRelative GuardianType = "Relative"
)

type Guardian struct {
	ID           uint
	StudentID    uint
	GuardianType string
	GuardianName string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}
