package service

import (
	"School_gql/config"
	"School_gql/pkg/model"
	"errors"

	"gorm.io/gorm"
)

type StaffService struct {
	db gorm.DB
}

func NewStaffService() *StaffService {
	staffService := new(StaffService)
	staffService.db = *config.GetDB()
	return staffService
}

func (f *StaffService) CreateStaff(staffService *model.Staff) (*model.Staff, error) {
	err := f.db.Where(&model.Staff{
		//StaffType: staffService.StaffType,
		Name:     staffService.Name,
		Aadharno: staffService.Aadharno,
	}).FirstOrCreate(&staffService).Error
	if err != nil {
		return nil, errors.New("can't create staff")
	}
	return staffService, nil
}

func (f *StaffService) DeleteStaff(staffID string) (bool, error) {
	err := f.db.Where("id=?", staffID).Delete(&model.Staff{}).Error
	if err != nil {
		return false, errors.New("can't delete staff")
	}
	return true, nil
}

func (f *StaffService) GetStaff(staffType string) ([]*model.Staff, error) {
	staffData := []*model.Staff{}
	err := f.db.Where("staff_type ILIKE ?", staffType).Find(&staffData).Error
	if err != nil {
		return nil, errors.New("can't get staff")
	}
	return staffData, nil
}
func (f *StaffService) GetStaffByID(staffType string, staffID string) (*model.Staff, error) {
	staffData := &model.Staff{}
	err := f.db.Where("staff_type ILIKE ?", staffType).Where("id = ?", staffID).Find(staffData).Error
	if err != nil {
		return nil, errors.New("can't get staff")
	}
	return staffData, nil
}
