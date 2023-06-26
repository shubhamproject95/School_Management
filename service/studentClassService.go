package service

import (
	"School_gql/config"
	"School_gql/pkg/model"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type StudentClassService struct {
	db *gorm.DB
}

func NewStudentClassService() *StudentClassService {
	studentClass := new(StudentClassService)
	studentClass.db = config.GetDB()
	return studentClass
}

func (s *StudentClassService) CreateStudentClassService(studentClass *model.StudentClass) (*model.StudentClass, error) {
	err := s.db.Where(model.StudentClass{
		StudentID:     studentClass.StudentID,
		ClassID:       studentClass.ClassID,
		SessionYearID: studentClass.SessionYearID,
	}).FirstOrCreate(&studentClass).Error

	// var perPage float64 = 20
	// var totalRecord float64
	// s.db.Raw("SELECT COUNT(id) FROM users").Scan(&totalRecord)
	// totalpage = math.Ceil(totalRecord / perPage)
	// nextOffset = (totalRecord - 1) * perPage

	if err != nil {
		return nil, errors.New("can't create student service")
	}
	return studentClass, nil
}

func (s *StudentClassService) DeleteStudentClassService(studentClassID string) (bool, error) {
	studentClass := &model.StudentClass{}
	err := s.db.Where("id =?", studentClassID).Delete(studentClass).Error
	if err != nil {
		return false, errors.New("can't delete student class")
	}
	return true, nil
}

func (s *StudentClassService) UpdateStudentClassService(studentClass *model.StudentClass) (*model.StudentClass, error) {
	err := s.db.Where("id=?", studentClass.ID).Updates(&studentClass).Error
	if err != nil {
		return nil, errors.New("can't update student class")
	}
	return studentClass, nil
}

func (s *StudentClassService) GetStudentClassService(studentClassID string) (bool, error) {
	studentClass := &model.StudentClass{}
	err := s.db.Model(&model.StudentClass{}).Preload(clause.Associations).First(studentClass, "id=?", studentClassID)
	if err != nil {
		return false, errors.New("can't get student details")
	}
	return true, nil
}

func (s *StudentClassService) GetsStudentClassService() (*[]model.StudentClass, error) {
	studentClass := &[]model.StudentClass{}
	err := s.db.Model(&[]model.StudentClass{}).Find(studentClass).Error
	if err != nil {
		return nil, errors.New("can't gets student")
	}
	return studentClass, nil
}
