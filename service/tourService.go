package service

import (
	"School_gql/config"
	"School_gql/pkg/model"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TourService struct {
	db gorm.DB
}

func NewTourService() *TourService {
	tourservice := new(TourService)
	tourservice.db = *config.GetDB()
	return tourservice
}

func (t *TourService) CreateTour(tour *model.Tour) (*model.Tour, error) {
	err := t.db.Where(model.Tour{
		StudentID: tour.StudentID,
		ClassID:   tour.ClassID,
		TeacherID: tour.TeacherID,
	}).FirstOrCreate(&tour).Error

	if err != nil {
		return nil, errors.New("can't create tours")
	}
	return tour, nil
}

func (t *TourService) DeleteTour(tourID string) (bool, error) {
	err := t.db.Where("id=?", tourID).Delete(&model.Tour{}).Error
	if err != nil {
		return false, errors.New("can't delete tours")
	}
	return true, nil
}

func (t *TourService) UpdateTour(tour *model.Tour) (bool, error) {
	err := t.db.Where("id=?", tour.ID).Updates(&model.Tour{}).Error
	if err != nil {
		return false, errors.New("can't update tour")
	}
	return true, nil
}

func (t *TourService) GetTour(tourID string) (*model.Tour, error) {
	tour := &model.Tour{}
	err := t.db.Model(&model.Tour{}).Preload(clause.Associations).First("id=?", tourID)
	if err != nil {
		return nil, errors.New("can't get tour")
	}
	return tour, nil
}

func (t *TourService) GetsTours() (*[]model.Tour, error) {
	tour := &[]model.Tour{}
	err := t.db.Model(&[]model.Tour{}).Find(tour).Error
	if err != nil {
		return nil, errors.New("can't gets student details")
	}
	return tour, nil
}
