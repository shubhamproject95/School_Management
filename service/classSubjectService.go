package service

import (
	"School_gql/config"
	"School_gql/pkg/model"
	"errors"

	"gorm.io/gorm"
)

type ClassSubject struct {
	db *gorm.DB
}

func NewClassSubjectService() *ClassSubject {
	classSubject := new(ClassSubject)
	classSubject.db = config.GetDB()
	return classSubject
}

func (s *ClassSubject) CreateClassSubject(classSubject *model.ClassSubject) (*model.ClassSubject, error) {
	err := s.db.Where(model.ClassSubject{
		ClassID:   classSubject.ClassID,
		SubjectID: classSubject.SubjectID,
	}).FirstOrCreate(&classSubject).Error
	if err != nil {
		return nil, err
	}
	return classSubject, nil
}

func (s *ClassSubject) DeleteClassSubject(ClassSubjectID int) (bool, error) {
	err := s.db.Where("id=?", ClassSubjectID).Delete(&model.ClassSubject{}).Error
	if err != nil {
		return false, errors.New("can't delete class")
	}
	return true, nil
}

func (s *ClassSubject) GetClassSubject(classSubjectID int) (*model.ClassSubject, error) {
	class := &model.ClassSubject{}
	err := s.db.Model(&model.ClassSubject{}).First(class, "id=?", classSubjectID).Error
	if err != nil {
		return nil, errors.New("can't get class subject")
	}
	return class, nil
}

func (s *ClassSubject) GetsClassSubject() (*[]model.ClassSubject, error) {
	subjects := &[]model.ClassSubject{}
	err := s.db.Model(subjects).Find(subjects).Error
	if err != nil {
		return nil, errors.New("can't get class subject")
	}
	return subjects, nil
}

func (s *ClassSubject) UpdateClassSubject(classSubject *model.ClassSubject) (bool, error) {
	err := s.db.Where("id=?", classSubject.ID).Updates(&classSubject).Error
	if err != nil {
		return false, errors.New("can't update class")
	}
	return true, nil
}
