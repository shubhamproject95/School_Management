package service

import (
	"School_gql/config"
	"School_gql/pkg/model"
	"errors"

	"gorm.io/gorm"
)

type GuardianService struct {
	db *gorm.DB
}

func NewGuardianService() *GuardianService {
	guardianService := new(GuardianService)
	guardianService.db = config.GetDB()
	return guardianService
}

func (g *GuardianService) CreateGuardian(guardian *model.Guardian) (*model.Guardian, error) {
	err := g.db.Where(model.Guardian{
		StudentID:    guardian.StudentID,
		GuardianName: guardian.GuardianName,
	}).FirstOrCreate(&guardian).Error
	if err != nil {
		return nil, errors.New("can't create guardian")
	}
	return guardian, nil
}

func (g *GuardianService) DeleteGuardian(guardianID string) (bool, error) {

	err := g.db.Where("id=?", guardianID).Delete(&model.Guardian{}).Error
	if err != nil {
		return false, errors.New("can't delete guardian")
	}
	return true, nil
}
