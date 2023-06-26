package service

import (
	"School_gql/config"
	"School_gql/pkg/model"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RoleService struct {
	db gorm.DB
}

func NewRoleervice() *RoleService {
	Roleervice := new(RoleService)
	Roleervice.db = *config.GetDB()
	return Roleervice
}

func (r *RoleService) CreateRole(role *model.Role) (*model.Role, error) {
	err := r.db.Where(model.Role{
		Role: role.Role,
	}).FirstOrCreate(&role).Error
	if err != nil {
		return nil, errors.New("can't create role")
	}
	return role, nil
}

func (r *RoleService) DeleteRole(roleID string) (bool, error) {
	err := r.db.Where("id=?", roleID).Delete(&model.Role{}).Error
	if err != nil {
		return false, errors.New("can't delete role")
	}
	return true, nil
}

func (r *RoleService) GetRole(roleID string) (*model.Role, error) {
	role := &model.Role{}
	err := r.db.Model(&model.Role{}).Preload(clause.Associations).First(role, "id=?", roleID).Error
	if err != nil {
		return nil, errors.New("can't get Role")
	}
	return role, nil
}

func (r *RoleService) UpdateRole(role *model.Role) (*model.Role, error) {
	err := r.db.Where("id=?", role.ID).Updates(&role).Error
	if err != nil {
		return nil, errors.New("can't update Role")
	}
	return role, nil
}
