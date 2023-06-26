package service

import (
	"School_gql/config"
	"School_gql/pkg/model"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type StudentService struct {
	db *gorm.DB
}

func NewStudentService() *StudentService {
	studentservice := new(StudentService)
	studentservice.db = config.GetDB()
	return studentservice
}

func (s *StudentService) CreateStudent(student *model.Student) (*model.Student, error) {
	err := s.db.Where(model.Student{
		Name:        student.Name,
		DOB:         student.DOB,
		Father_Name: student.Father_Name,
		Mother_Name: student.Mother_Name,
		Address:     student.Address,
		Status:      student.Status,
	}).FirstOrCreate(&student).Error
	if err != nil {
		return nil, err
	}
	return student, nil
}

func (s *StudentService) DeleteStudent(studentID string) (bool, error) {
	err := s.db.Where("id = ?", studentID).Delete(&model.Student{}).Error
	if err != nil {
		return false, errors.New("can't delete student")
	}
	return true, nil
}

func (s *StudentService) GetStudent(studentID string) (*model.Student, error) {
	student := &model.Student{}
	err := s.db.
		Model(&model.Student{}).
		Preload(clause.Associations).
		First(student, "id=?", studentID).Error

	if err != nil {
		return nil, errors.New("can't get student")
	}
	return student, nil
}

func (s *StudentService) GetsStudent() ([]*model.Student, error) {
	students := []*model.Student{}
	err := s.db.Model(students).Find(&students).Error
	if err != nil {
		return nil, errors.New("can't get students")
	}
	return students, nil
}

func (s *StudentService) UpdateStudent(studentService *model.Student) (bool, error) {
	err := s.db.Where("id=?", studentService.ID).Updates(&studentService).Error
	if err != nil {
		return false, errors.New("can't update student service")
	}
	return true, nil
}
