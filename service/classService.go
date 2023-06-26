package service

import (
	"School_gql/config"
	"School_gql/pkg/model"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ClassService struct {
	db gorm.DB
}

func NewClassService() *ClassService {
	classService := new(ClassService)
	classService.db = *config.GetDB()
	return classService
}

func (c *ClassService) CreateClass(class *model.Class) (*model.Class, error) {
	err := c.db.Where(model.Class{
		Class: class.Class,
	}).FirstOrCreate(&class).Error
	if err != nil {
		return nil, errors.New("can't create a class")
	}
	return class, nil
}

func (c *ClassService) DeleteClass(classID string) (bool, error) {
	err := c.db.Where("id=?", classID).Delete(&model.Class{}).Error
	if err != nil {
		return false, errors.New("can't delete student")
	}
	return true, nil
}

func (c *ClassService) UpdateClass(class *model.Class) (bool, error) {
	err := c.db.Where("id=?", class.ID).Updates(&class).Error
	if err != nil {
		return false, errors.New("can't update class")
	}
	return true, nil
}

func (c *ClassService) GetClass(classID string) (*model.Class, error) {
	class := &model.Class{}
	err := c.db.Model(&model.Class{}).Preload(clause.Associations).First(class, "id=?", classID).Error
	if err != nil {
		return nil, errors.New("can't get class ID")
	}
	return class, nil
}

func (c *ClassService) GetClasses() (*[]model.Class, error) {
	classes := &[]model.Class{}
	err := c.db.Model(&[]model.Class{}).Find(classes).Error
	if err != nil {
		return nil, errors.New("can't get classes")
	}
	return classes, nil
}
