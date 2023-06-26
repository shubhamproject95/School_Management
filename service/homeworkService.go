package service

import (
	"School_gql/config"
	"School_gql/pkg/model"
	"errors"

	"gorm.io/gorm"
)

type HomeworkService struct {
	db gorm.DB
}

func NewHomeworkService() *HomeworkService {
	homeworkService := new(HomeworkService)
	homeworkService.db = *config.GetDB()
	return homeworkService
}

func (h *HomeworkService) CreateHomework(homework *model.Homework) (*model.Homework, error) {

	hm := &homework
	err := h.db.Debug().Where(hm).FirstOrCreate(&homework).Error
	if err != nil {
		return nil, errors.New("can't create homework")
	}
	return homework, nil
}

func (h *HomeworkService) DeleteHomework(homeworkID string) (bool, error) {
	err := h.db.Where("id=?", homeworkID).Delete(&model.Homework{}).Error
	if err != nil {
		return false, errors.New("can't delete homework")
	}
	return true, nil
}
