package service

import (
	"School_gql/config"
	"School_gql/pkg/model"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type FeesService struct {
	db *gorm.DB
}

func NewFeeService() *FeesService {
	feesService := new(FeesService)
	feesService.db = config.GetDB()
	return feesService
}

func (f *FeesService) CreateFees(fees *model.Fees) (*model.Fees, error) {

	err := f.db.
		Where(model.Fees{
			ClassID:   fees.ClassID,
			StudentID: fees.StudentID,
			Session:   fees.Session,
			FeesPaid:  fees.FeesPaid,
			TotalFees: fees.TotalFees,
			Pending:   fees.Pending,
			Month:     fees.Month,
		}).
		FirstOrCreate(fees).Error

	if err != nil {
		return nil, errors.New("can't create fees")
	}
	return fees, nil
}

func (f *FeesService) DeleteFees(feesID string) (bool, error) {
	fees := &model.Fees{}
	err := f.db.Where("id=?", feesID).Delete(fees).Error
	if err != nil {
		return false, errors.New("can't delete fees")
	}
	return true, nil
}

func (f *FeesService) UpdateFees(fees *model.Fees) (bool, error) {
	err := f.db.Where("id=", fees.ID).Updates(&fees)
	if err != nil {
		return false, errors.New("can't update fees")
	}
	return true, nil
}

func (f *FeesService) GetFees(feesID string) (*model.Fees, error) {
	fees := &model.Fees{}
	err := f.db.Model(&model.Fees{}).Preload(clause.Associations).First(fees, "id=?", feesID)
	if err != nil {
		return nil, err.Error
	}
	return fees, nil
}

func (f *FeesService) GetsFees() (*[]model.Fees, error) {
	fees := &[]model.Fees{}
	err := f.db.Model(&[]model.Fees{}).Find(fees).Error
	if err != nil {
		return nil, errors.New("can't gets fees")
	}
	return fees, nil
}
