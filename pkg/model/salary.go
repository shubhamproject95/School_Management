package model

import "time"

type Salary struct {
	ID           uint   `gorm:"autoIncrement;primary_key"`
	Basic        uint   `gorm:"not null;default:0"`
	Bonus        uint   `gorm:"not null;default:0"`
	PF           uint   `gorm:"not null;default:0"`
	Tax          uint   `gorm:"not null;default:0"`
	NetAmount    uint   `gorm:"not null;default:0"`
	Month        string `gorm:"not null;default:0"`
	StaffID      uint   `gorm:"foreignKey"`
	BankDetailID uint   `gorm:"foreignKey"`
	CreatedAt    time.Time
	DeletedAt    time.Time
	UpdatedAt    time.Time
}
