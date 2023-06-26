package service

import (
	"School_gql/config"
	"School_gql/pkg/model"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ClassTeacherService struct {
	db *gorm.DB
}

func NewClassTeacherService() *ClassTeacherService {
	classTeacherService := new(ClassTeacherService)
	classTeacherService.db = config.GetDB()
	return classTeacherService
}

func (t *ClassTeacherService) CreateClassTeacher(classTeacher *model.ClassTeacher) (*model.ClassTeacher, error) {
	err := t.db.Where(model.ClassTeacher{
		ClassID:   classTeacher.ClassID,
		TeacherID: classTeacher.TeacherID,
	}).FirstOrCreate(&classTeacher).Error
	if err != nil {
		return nil, errors.New("can't Create Class Teacher")
	}
	return classTeacher, nil
}

func (t *ClassTeacherService) DeleteClassTeacher(classTeacherID string) (bool, error) {
	classTeacher := &model.ClassTeacher{}
	err := t.db.Where("id = ?", classTeacherID).Delete(classTeacher).Error
	if err != nil {
		return false, errors.New("can't delete Class Teacher")
	}
	return true, nil
}

func (t *ClassTeacherService) UpdateClassTeacher(classTeacher *model.ClassTeacher) (bool, error) {
	err := t.db.Where("id =?", classTeacher.ID).Updates(&classTeacher).Error
	if err != nil {
		return false, errors.New("can't update class teacher")
	}
	return true, nil
}

func (t *ClassTeacherService) GetClassTeacher(classTeacherID string) (bool, error) {
	classTeacher := &model.ClassTeacher{}
	err := t.db.Model(&model.ClassTeacher{}).Preload(clause.Associations).First(classTeacher, "id=?", classTeacherID).Error
	if err != nil {
		return false, errors.New("can't get class ID")
	}
	return true, nil
}

func (t *ClassTeacherService) GetsClassTeacher() (*[]model.ClassTeacher, error) {
	classTeacher := &[]model.ClassTeacher{}
	err := t.db.Model(&[]model.ClassTeacher{}).Find(classTeacher).Error
	if err != nil {
		return nil, errors.New("can't get class teacher")
	}
	return classTeacher, nil
}
