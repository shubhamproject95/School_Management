package service

import (
	"School_gql/config"
	"School_gql/pkg/model"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SubjectService struct {
	db *gorm.DB
}

func NewSubjectService() *SubjectService {
	subject := new(SubjectService)
	subject.db = config.GetDB()
	return subject
}

func (s *SubjectService) CreateSubject(subject *model.Subject) (*model.Subject, error) {
	err := s.db.Where(&model.Subject{

		Subject: subject.Subject,
	}).FirstOrCreate(&subject).Error

	if err != nil {
		return nil, errors.New("can't create subject")
	}
	return subject, nil
}

func (s *SubjectService) DeleteSubject(subjectID string) (bool, error) {
	subject := &model.Subject{}
	err := s.db.Where("id=?", subjectID).Delete(subject).Error
	if err != nil {
		return false, errors.New("can't delete subject")
	}
	return true, nil
}

func (s *SubjectService) UpdateSubject(subject *model.Subject) (*model.Subject, error) {
	//subjectobj := &model.Subject{}
	err := s.db.Where("id=?", subject.ID).Updates(&subject).Error
	if err != nil {
		return nil, errors.New("can't update subject")
	}
	return subject, nil
}

func (s *SubjectService) GetSubject(subjectID string) (*model.Subject, error) {
	subject := &model.Subject{}
	err := s.db.Model(&model.Subject{}).Preload(clause.Associations).First(subject, "id=?", subjectID).Error
	if err != nil {
		return nil, errors.New("can't get subject id")
	}
	return subject, nil
}

func (s *SubjectService) GetsSubject() (*[]model.Subject, error) {
	subject := &[]model.Subject{}
	err := s.db.Model(&[]model.Subject{}).Find(subject).Error
	if err != nil {
		return nil, errors.New("can't get subjects")
	}
	return subject, nil
}
