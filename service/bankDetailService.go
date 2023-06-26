package service

import (
	"School_gql/config"
	"School_gql/pkg/model"
	"errors"

	"gorm.io/gorm"
)

type BankDetail struct {
	db *gorm.DB
}

func NewBankDetailervice() *BankDetail {
	BankDetail := new(BankDetail)
	BankDetail.db = config.GetDB()
	return BankDetail
}

// Create bank details
func (b *BankDetail) CreateBankDetail(bankDetail *model.BankDetail) (*model.BankDetail, error) {
	err := b.db.Where(model.BankDetail{
		Bank:        bankDetail.Bank,
		BankAccount: bankDetail.BankAccount,
		Name:        bankDetail.Name,
	}).FirstOrCreate(&bankDetail).Error
	if err != nil {
		return nil, errors.New("can't enter bank details")
	}
	return bankDetail, nil
}

// Delete bank details
func (b *BankDetail) DeleteBankDetail(BankDetailID string) (bool, error) {
	err := b.db.Where("id=?", BankDetailID).Delete(&model.BankDetail{}).Error
	if err != nil {
		return false, nil
	}
	return true, nil
}

// Get Bank details
func (b *BankDetail) GetsBankDetail(staffID uint) (*[]model.BankDetail, error) {
	BankDetail := &[]model.BankDetail{}
	err := b.db.Where("staff_id=?", staffID).Find(BankDetail).Error
	if err != nil {
		return nil, errors.New("can't get bank details")
	}
	return BankDetail, nil
}
