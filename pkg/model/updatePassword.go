package model

type UpdatePassword struct {
	ID              uint `gorm:"primary_key:autoIncrement"`
	UserID          uint `gorm:"foreignKey"`
	Password        string
	ConformPassword string
	ResetPassword   ResetPassword
}
