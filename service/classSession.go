package service

import (
	"School_gql/config"
	"School_gql/pkg/model"

	"errors"

	"gorm.io/gorm"
)

type ClassSessionService struct {
	db *gorm.DB
}

func NewClassSessionService() *ClassSessionService {
	classService := new(ClassSessionService)
	classService.db = config.GetDB()
	return classService
}

func (c *ClassSessionService) TotalClassStudents(classID int, sessionID int) (float64, error) {
	var count int64
	err := c.db.Model(&model.StudentClass{}).Where("class_id=? AND session_year_id=?", classID, sessionID).Count(&count).Error
	if err != nil {
		return 0, nil
	}
	return float64(count), nil
}

func (c *ClassSessionService) GetClassStudents(classID int, sessionID int, offset int) (*[]model.StudentClass, error) {

	stuclass := &[]model.StudentClass{}

	err := c.db.Model(&model.StudentClass{}).Where("class_id=? AND session_year_id=?", classID, sessionID).Limit(20).Offset(offset).Find(stuclass).Error
	if err != nil {
		return nil, errors.New("can't get total pages")
	}
	return stuclass, nil
}
