package service

import (
	"School_gql/config"
	"School_gql/pkg/model"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type StudentAttendence struct {
	db *gorm.DB
}

func NewAttendenceService() *StudentAttendence {
	studentAttendence := new(StudentAttendence)
	studentAttendence.db = config.GetDB()
	return studentAttendence
}

func (a *StudentAttendence) CreateAttendence(attendence *model.StudentAttendence) (*model.StudentAttendence, error) {
	err := a.db.Where(&model.StudentAttendence{
		StudentID:  attendence.StudentID,
		Attendence: attendence.Attendence,
		Date:       attendence.Date,
	}).FirstOrCreate(&attendence).Error
	if err != nil {
		return nil, errors.New("can't create attendence")
	}
	return attendence, nil
}

func (a *StudentAttendence) DeleteAttendence(attendenceID string) (bool, error) {
	attendence := &model.StudentAttendence{}
	err := a.db.Where("id=?", attendenceID).Delete(attendence).Error
	if err != nil {
		return false, errors.New("can't delete attendence")
	}
	return true, nil
}

func (a *StudentAttendence) UpdateAttendence(attendence *model.StudentAttendence) (*model.StudentAttendence, error) {
	//stuattendence := &model.StudentAttendence{}
	err := a.db.Where("id=?", attendence.ID).Updates(&attendence).Error
	if err != nil {
		return nil, errors.New("can't update attendence")
	}
	return attendence, nil
}

func (a *StudentAttendence) GetAttendence(attendenceID string) (*model.StudentAttendence, error) {
	attendence := &model.StudentAttendence{}
	err := a.db.Model(&model.StudentAttendence{}).Preload(clause.Associations).First(attendence, "id=?", attendenceID)
	if err != nil {
		return nil, errors.New("can't get id")
	}
	return attendence, nil
}

func (a *StudentAttendence) GetsAttendence() (*[]model.StudentAttendence, error) {
	attendence := &[]model.StudentAttendence{}
	err := a.db.Model(&[]model.StudentAttendence{}).Find(attendence).Error
	if err != nil {
		return nil, errors.New("can't get attendence")
	}
	return attendence, nil
}
