package service

import (
	"School_gql/config"
	"School_gql/pkg"
	"School_gql/pkg/model"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserService struct {
	db gorm.DB
}

func NewUserService() *UserService {
	UserService := new(UserService)
	UserService.db = *config.GetDB()
	return UserService
}

func (l *UserService) CreateUser(user *model.User) (*model.User, error) {
	err := l.db.Where(model.User{
		UserName: user.UserName,
		Password: user.Password,
		UserType: user.UserType,
	}).FirstOrCreate(&user).Error

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (l *UserService) DeleteUser(userID string) (bool, error) {
	err := l.db.Where("id=?", userID).Delete(&model.User{}).Error
	if err != nil {
		return false, errors.New("can't delete User")
	}
	return true, nil
}

func (l *UserService) GetUser(userID string) (*model.User, error) {
	user := &model.User{}
	err := l.db.Model(&model.User{}).Preload(clause.Associations).First(user, "id=?", userID).Error
	if err != nil {
		return nil, errors.New("can't get User")
	}
	return user, nil
}

func (l *UserService) UpdateUser(user *model.User) (*model.User, error) {
	err := l.db.Where("id=?", user.ID).Updates(&user).Error
	if err != nil {
		return nil, errors.New("can't update User details")
	}
	return user, nil
}

func (l *UserService) UserLogin(username string, password string) (*model.Login, error) {
	userLogin := &model.Login{}
	err := l.db.Where("user_name=? AND Password=?", username, password).Preload("Role").First(userLogin).Error
	if err != nil {
		return nil, errors.New("invalid user")
	}
	return userLogin, nil
}

func (l *UserService) GetUserByModel(user *model.User) (*model.User, error) {
	userData := model.User{}
	err := l.db.Where(&user).First(&userData).Error
	if err != nil {
		return nil, errors.New("can't get password")
	}
	return &userData, nil
}

func (l *UserService) GetResetKey(resetKey string) (*model.ResetPassword, error) {
	resetPassword := model.ResetPassword{}
	err := l.db.Where(&model.ResetPassword{
		ResetKey: resetKey,
		IsUsed:   false,
	}).First(&resetPassword).Error
	if err != nil {
		return nil, err
	}
	return &resetPassword, nil
}

func (l *UserService) UpdateResetKey(resetPassword *model.ResetPassword) (bool, error) {
	err := l.db.Where(resetPassword).Updates(&model.ResetPassword{

		IsUsed: true,
	}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (l *UserService) Auth(token string) (*model.Login, error) {
	user := model.Login{}
	if err := l.db.Where("token=?", token).
		Preload("Role").
		First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (l *UserService) GenerateToken(loginID uint) (string, error) {
	token := pkg.RandomString()
	err := l.db.Where("id=?", loginID).Updates(&model.Login{Token: token}).Error
	if err != nil {
		return "", err
	}
	return token, nil
}
