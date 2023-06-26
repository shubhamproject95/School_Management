package service

import (
	"School_gql/config"
	"School_gql/pkg/model"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SessionYear struct {
	db *gorm.DB
}

func NewSessionYear() *SessionYear {
	SessionYear := new(SessionYear)
	SessionYear.db = config.GetDB()
	return SessionYear
}

func (s *SessionYear) CreateSessionYear(session *model.SessionYear) (*model.SessionYear, error) {
	err := s.db.Where(model.SessionYear{
		SessionYear: session.SessionYear,
	}).FirstOrCreate(&session).Error
	if err != nil {
		return nil, errors.New("can't create student seassion")
	}
	return session, nil
}

func (s *SessionYear) DeleteSessionYear(sessionID string) (bool, error) {
	session := &model.SessionYear{}
	err := s.db.Where("id=?", sessionID).Delete(session).Error
	if err != nil {
		return false, errors.New("can't delete session")
	}
	return true, nil
}

func (s *SessionYear) UpdateSessionYear(session *model.SessionYear) (bool, error) {
	err := s.db.Where("id=", session.ID).First(&session).Error
	if err != nil {
		return false, errors.New("can't update session")
	}
	return true, nil
}

func (s *SessionYear) GetSessionYear(sessionID string) (*model.SessionYear, error) {
	session := &model.SessionYear{}
	err := s.db.Model(&model.SessionYear{}).Preload(clause.Associations).First(session, "id=?", sessionID).Error
	if err != nil {
		return nil, errors.New("can't get seassion Id")
	}
	return session, nil
}

func (s *SessionYear) GetsSessionYear() (*[]model.SessionYear, error) {
	session := &[]model.SessionYear{}
	err := s.db.Model(&[]model.SessionYear{}).Find(session).Error
	if err != nil {
		return nil, errors.New("can't gert student session")
	}
	return session, nil
}
