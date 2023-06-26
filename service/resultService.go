package service

import (
	"School_gql/config"
	"School_gql/pkg/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ResultService struct {
	db *gorm.DB
}

func NewResultService() *ResultService {
	resultService := new(ResultService)
	resultService.db = config.GetDB()
	return resultService
}

func (r *ResultService) CreateResult(result *model.Result) (*model.Result, error) {
	err := r.db.Where(model.Result{
		ResultType: result.ResultType,
		RollNumber: result.RollNumber,
	}).FirstOrCreate(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *ResultService) DeleteResult(resultID string) (bool, error) {
	err := r.db.Where("id=?", resultID).Delete(model.Result{}).Error
	if err != nil {
		return false, nil
	}
	return true, nil
}

func (r *ResultService) GetResult(resultID string) (*model.Result, error) {
	result := &model.Result{}
	err := r.db.Model(&model.Result{}).Preload(clause.Associations).First(result, "id=?", resultID).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *ResultService) UpdateResult(result *model.Result) (*model.Result, error) {
	err := r.db.Where("id=?", result.ID).Updates(&model.Result{}).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
