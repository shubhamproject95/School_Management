package model

type Scholarship struct {
	ID              uint `gorm:"primary_key:autoIncrement"`
	StudentID       uint `gorm:"foreignKey"`
	ClassID         uint `gorm:"foreignKey"`
	StudentResultID uint `gorm:"foreignKey"`
	CreatedAt       uint
	DeletedAt       uint
	UpdatedAt       uint
}
