package service

import (
	"School_gql/config"
	"School_gql/pkg/model"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ScholarshipService struct {
	db gorm.DB
}

func NewScholarshipService() *ScholarshipService {
	scholarshipService := new(ScholarshipService)
	scholarshipService.db = *config.GetDB()
	return scholarshipService
}

func (s *ScholarshipService) CreateSchoarShip(scholarShip *model.Scholarship) (*model.Scholarship, error) {
	err := s.db.Where(&model.Scholarship{
		StudentID:       scholarShip.StudentID,
		ClassID:         scholarShip.ClassID,
		StudentResultID: scholarShip.StudentResultID,
	}).FirstOrCreate(&scholarShip).Error
	if err != nil {
		return nil, errors.New("CAn't create scholarship")
	}
	return scholarShip, nil
}

func (s *ScholarshipService) DeleteScholarship(scholarShipID string) (bool, error) {
	err := s.db.Where("id=?", scholarShipID).Error
	if err != nil {
		return false, errors.New("can't deelte Scholarship")
	}
	return true, nil
}

func (s *ScholarshipService) GetScholarshipStudent(scholarShipID string) (*model.Scholarship, error) {
	scholarsip := &model.Scholarship{}
	err := s.db.Model(&model.Scholarship{}).Preload(clause.Associations).First(scholarShipID, "id=?", scholarShipID).Error
	if err != nil {
		return nil, errors.New("can't get scholarship students")
	}
	return scholarsip, nil

}

func (s *ScholarshipService) GetsScholarshipStudents() (*[]model.Scholarship, error) {
	scholarship := &[]model.Scholarship{}
	err := s.db.Model(&[]model.Scholarship{}).Find(scholarship).Error
	if err != nil {
		return nil, errors.New("can't get student scholarships")
	}
	return scholarship, nil
}

func (s *ScholarshipService) UpdateScholarship(scholarship *model.Scholarship) (bool, error) {
	err := s.db.Where("id=?", scholarship.ID).Updates(&model.Scholarship{}).Error
	if err != nil {
		return false, errors.New("can't update scholarship")
	}
	return true, nil

}
