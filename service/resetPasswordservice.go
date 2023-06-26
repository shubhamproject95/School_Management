package service

import (
	"School_gql/config"
	"School_gql/pkg/model"
	"errors"

	"gorm.io/gorm"
)

type ResetPasswordService struct {
	db gorm.DB
}

func NewResetPasswordService() *ResetPasswordService {
	resetPasswordService := new(ResetPasswordService)
	resetPasswordService.db = *config.GetDB()
	return resetPasswordService
}

func (r *ResetPasswordService) SaveRestKey(resetPassword *model.ResetPassword) (*model.ResetPassword, error) {
	err := r.db.Create(&resetPassword).Error

	if err != nil {
		return nil, errors.New("can't reset password")
	}
	return resetPassword, nil
}

func (r *ResetPasswordService) GetRestKey(resetKey *model.ResetPassword) (*model.ResetPassword, error) {
	err := r.db.Where("id=?", resetKey.ID).First(&resetKey).Error
	if err != nil {
		return nil, errors.New("can't get reset key")
	}
	return resetKey, nil
}
