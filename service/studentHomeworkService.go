package service

import (
	"School_gql/config"
	"School_gql/pkg/model"
	"errors"

	"gorm.io/gorm"
)

type StudentHomework struct {
	db *gorm.DB
}

func NewStudentHomeWorkService() *StudentHomework {
	studentHomework := new(StudentHomework)
	studentHomework.db = config.GetDB()
	return studentHomework
}

func (s *StudentHomework) CreateStudentHomework(studentHomework *model.StudentHomework) (*model.StudentHomework, error) {
	err := s.db.Where(model.StudentHomework{
		HomeworkID: studentHomework.HomeworkID,
		StudentID:  studentHomework.StudentID,
	}).FirstOrCreate(&studentHomework).Error
	if err != nil {
		return nil, errors.New("can't create student homework")
	}
	return studentHomework, nil
}

func (s *StudentHomework) DeleteStudentHomework(studentHomeworkID string) (bool, error) {
	err := s.db.Where("id=?", studentHomeworkID).Delete(&model.StudentHomework{}).Error
	if err != nil {
		return false, errors.New("can't delete student homework")
	}
	return true, nil
}
