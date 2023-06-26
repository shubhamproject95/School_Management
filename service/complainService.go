package service

import (
	"School_gql/config"
	"School_gql/pkg/model"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ComplainService struct {
	db *gorm.DB
}

func NewComplainService() *ComplainService {
	complainService := new(ComplainService)
	complainService.db = config.GetDB()
	return complainService
}

func (c *ComplainService) CreateComplain(complain *model.Complain) (*model.Complain, error) {

	err := c.db.Where(&model.Complain{
		StudentID: complain.StudentID,
		Complain:  complain.Complain,
		Student: model.Student{
			ID: complain.StudentID,
		},
	}).FirstOrCreate(&complain).Error

	if err != nil {
		return nil, errors.New("can't create complain")
	}
	return complain, nil
}

func (c *ComplainService) DeleteComplain(complainID string) (bool, error) {
	complain := &model.Complain{}
	err := c.db.Where("id =?", complainID).Delete(complain).Error
	if err != nil {
		return false, errors.New("can't delete complain")
	}
	return true, nil
}

func (c *ComplainService) GetComplain(complainID string) (*model.Complain, error) {
	complain := &model.Complain{}
	err := c.db.Model(&model.Complain{}).Preload(clause.Associations).First(complain, "id=?", complainID).Error
	if err != nil {
		return nil, errors.New("can't get complain")
	}
	return complain, nil
}

func (c *ComplainService) GetsComplain() (*[]model.Complain, error) {
	complain := &[]model.Complain{}
	err := c.db.Model(&[]model.Complain{}).Find(complain).Error
	if err != nil {
		return nil, errors.New("can't get complain")
	}
	return complain, nil
}

func (c *ComplainService) UpdateComplain(complain *model.Complain) (*model.Complain, error) {
	err := c.db.Where("id=?", complain.ID).Updates(&complain).Error
	if err != nil {
		return nil, errors.New("can't update complain")
	}
	return complain, nil
}
