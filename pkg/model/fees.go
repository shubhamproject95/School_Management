package model

import (
	"time"
)

type Fees struct {
	ID        uint   `gorm:"primary_key:autoIncrement"`
	StudentID uint   `gorm:"foreignKey"`
	Session   string `gorm:"not null;default:0"`
	FeesPaid  uint   `gorm:"not null;default:0;faker:intinrange(5000,12000)"`
	TotalFees uint   `gorm:"not null;default:0"`
	Pending   uint   `gorm:"not null;default:0"`
	Month     string
	ClassID   uint      `gorm:"foreignKey"`
	CreatedAt time.Time `gorm:"autoUpdate"`
	UpdatedAt time.Time `gorm:"autoUpdate"`
}
