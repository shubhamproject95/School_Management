package model

import "time"

type ResetPassword struct {
	ID               uint `gorm:"primary_key:autoIncrement"`
	UserID           uint `gorm:"foreignKey"`
	ResetKey         string
	ValidTill        time.Time `gorm:"autoIncrement;default:null"`
	IsUsed           bool
	UpdatePasswordID uint
}
