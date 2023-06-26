package service

import (
	"School_gql/config"
	"School_gql/pkg/model"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TeacherService struct {
	db *gorm.DB
}

func NewTeacherService() *TeacherService {
	teacher := new(TeacherService)
	teacher.db = config.GetDB()
	return teacher
}

func (t *TeacherService) CreateTeacher(teacher *model.Teacher) (*model.Teacher, error) {
	err := t.db.Where(model.Teacher{
		FirstName:  teacher.FirstName,
		LastName:   teacher.LastName,
		Department: teacher.Department,
		DOB:        teacher.DOB,
		JoiningAt:  teacher.JoiningAt,
		Status:     teacher.Status,
	}).FirstOrCreate(&teacher).Error
	if err != nil {
		return nil, errors.New("can't create teacher")
	}
	return teacher, nil
}

func (t *TeacherService) DeleteTeacher(teacherID string) (bool, error) {
	teacher := &model.Teacher{}
	err := t.db.Where("id=?", teacherID).Delete(teacher).Error
	if err != nil {
		return false, errors.New("can't delete teacher")
	}
	return true, nil
}

func (t *TeacherService) UpdateTeacher(teacher *model.Teacher) (*model.Teacher, error) {
	err := t.db.Where("id=?", teacher.ID).Updates(&teacher).Error
	if err != nil {
		return nil, errors.New("can't update teacher")
	}
	return teacher, nil
}

func (t *TeacherService) GetTeacher(teacherID string) (*model.Teacher, error) {
	teacher := &model.Teacher{}
	err := t.db.Model(&model.Teacher{}).Preload(clause.Associations).First(teacher, "id=?", teacherID).Error
	if err != nil {
		return nil, errors.New("can't get teacher")
	}
	return teacher, nil
}

func (t *TeacherService) GetsTeacher() (*[]model.Teacher, error) {
	teacher := &[]model.Teacher{}
	err := t.db.Model(&[]model.Teacher{}).Find(teacher).Error
	if err != nil {
		return nil, errors.New("can't get teachers details")
	}
	return teacher, nil
}
