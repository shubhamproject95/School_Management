package service

import (
	"School_gql/config"
	"School_gql/pkg/model"

	"gorm.io/gorm"
)

type LoginService struct {
	db gorm.DB
}

func NewLoginService() *LoginService {
	loginservice := new(LoginService)
	loginservice.db = *config.GetDB()
	return loginservice
}

func (l *LoginService) CreateLogin(login *model.Login) (*model.Login, error) {
	err := l.db.Where(&model.Login{
		UserName:   login.UserName,
		Password:   login.Password,
		UserTypeID: login.UserTypeID,
		RoleID:     login.RoleID,
	}).FirstOrCreate(&login).Error
	if err != nil {
		return nil, err
	}
	return login, nil
}
