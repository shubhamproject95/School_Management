package service

import (
	"School_gql/config"
	"School_gql/pkg/model"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ExamService struct {
	db *gorm.DB
}

func NewExamService() *ExamService {
	examService := new(ExamService)
	examService.db = config.GetDB()
	return examService
}

func (e *ExamService) CreateExam(exam *model.Exam) (*model.Exam, error) {
	err := e.db.Where(model.Exam{
		SubjectID: exam.SubjectID,
		ClassID:   exam.ClassID,
	}).FirstOrCreate(&exam).Error
	if err != nil {
		return nil, errors.New("can't create exam")
	}
	return exam, nil
}

func (e *ExamService) DeleteExam(examID string) (bool, error) {
	exam := &model.Exam{}
	err := e.db.Where("id =?", examID).Delete(exam).Error
	if err != nil {
		return false, errors.New("can't delete exam")
	}
	return true, nil
}

func (e *ExamService) GetExam(examID string) (*model.Exam, error) {
	exam := &model.Exam{}
	err := e.db.Model(&model.Exam{}).Preload(clause.Associations).First(exam, "id=?", examID).Error
	if err != nil {
		return nil, errors.New("can't get exam")
	}
	return exam, nil
}

func (e *ExamService) GetsExam() (*[]model.Exam, error) {
	exam := &[]model.Exam{}
	err := e.db.Model(&model.Exam{}).Find(exam).Error
	if err != nil {
		return nil, errors.New("can't get exam id")
	}
	return exam, nil
}

func (e *ExamService) UpdateExam(exam *model.Exam) (*model.Exam, error) {
	err := e.db.Where("id=?", exam.ID).Updates(&exam).Error
	if err != nil {
		return nil, err
	}
	return exam, nil
}
