package service

import (
	"School_gql/config"
	"School_gql/pkg/model"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type StudentExamService struct {
	db *gorm.DB
}

func NewStudentExamService() *StudentExamService {
	examService := new(StudentExamService)
	examService.db = config.GetDB()
	return examService
}

func (e *StudentExamService) SaveStudentExamInfo(exam *model.StudentExam) (*model.StudentExam, error) {
	err := e.db.
		Where(&model.StudentExam{
			StudentID:     exam.StudentID,
			ExamID:        exam.ExamID,
			SessionYearID: exam.SessionYearID,
		}).
		FirstOrCreate(&exam).Error
	if err != nil {
		return nil, errors.New("can't create exam")
	}
	return exam, nil
}

func (e *StudentExamService) DeleteStudentExam(examID string) (bool, error) {
	exam := &model.StudentExam{}
	err := e.db.Where("id=?", examID).Delete(exam).Error
	if err != nil {
		return false, errors.New("can't delete exam")
	}
	return true, nil
}

func (e *StudentExamService) GetStudentExam(examID string) (*model.StudentExam, error) {
	exam := &model.StudentExam{}
	err := e.db.Model(&model.StudentExam{}).Preload(clause.Associations).First(exam, "id=?", examID).Error
	if err != nil {
		return nil, errors.New("can get Exam ID")
	}
	return exam, nil
}

func (e *StudentExamService) GetsStudentExam() (*[]model.StudentExam, error) {
	exam := &[]model.StudentExam{}
	err := e.db.Model(&[]model.StudentExam{}).Find(exam).Error
	if err != nil {
		return nil, errors.New("can't get studentexam")
	}
	return exam, nil
}

func (e *StudentExamService) UpdateStudentExam(studentExam *model.StudentExam) (bool, error) {
	err := e.db.Where("id=?", studentExam.ID).Updates(&studentExam)
	if err != nil {
		return false, errors.New("can't update student exam")
	}
	return true, nil
}

func (e *StudentExamService) GetStudentExamByID(studentID string) (*[]model.StudentExam, error) {
	exam := &[]model.StudentExam{}
	err := e.db.Where("student_id=?", studentID).Preload(clause.Associations).Find(exam).Error
	if err != nil {
		return nil, errors.New("can't get student exam details")
	}
	return exam, nil
}
