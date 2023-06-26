package service

import (
	"School_gql/config"
	"School_gql/pkg/model"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TeacherAttendence struct {
	db *gorm.DB
}

func NewTeacherAttendence() *TeacherAttendence {
	teacherAttendence := new(TeacherAttendence)
	teacherAttendence.db = config.GetDB()
	return teacherAttendence
}

func (t *TeacherAttendence) CreateTeacherAttendence(teacherAttendence *model.TeacherAttendence) (*model.TeacherAttendence, error) {
	err := t.db.Where(&model.TeacherAttendence{
		TeacherID:         teacherAttendence.TeacherID,
		TeacherAttendence: teacherAttendence.TeacherAttendence,
		Date:              teacherAttendence.Date,
	}).FirstOrCreate(&teacherAttendence).Error
	if err != nil {
		return nil, errors.New("can't create teacher attendence")
	}
	return teacherAttendence, nil
}

func (t *TeacherAttendence) DeleteTeacherAttendence(teacherAttendenceID string) (bool, error) {
	teacherAttendence := &model.TeacherAttendence{}
	err := t.db.Where("id=?", teacherAttendenceID).Delete(teacherAttendence).Error
	if err != nil {
		return false, errors.New("can't delete teacherAttendence")
	}
	return true, nil
}

func (t *TeacherAttendence) UpdateTeacherAttendence(teacherAttendence *model.TeacherAttendence) (*model.TeacherAttendence, error) {
	err := t.db.Where("id=?", teacherAttendence.ID).Updates(&teacherAttendence).Error
	if err != nil {
		return nil, errors.New("can't update teacher attendence")
	}
	return teacherAttendence, nil
}

func (t *TeacherAttendence) GetTeacherAttendence(teacherAttendenceID string) (*model.TeacherAttendence, error) {
	teacherAttendence := &model.TeacherAttendence{}
	err := t.db.Model(&model.TeacherAttendence{}).Preload(clause.Associations).First(teacherAttendence, "id=?", teacherAttendenceID).Error
	if err != nil {
		return nil, errors.New("can't get teacher attendence")
	}
	return teacherAttendence, nil
}

func (t *TeacherAttendence) GetsTeacherAttendence() (*[]model.TeacherAttendence, error) {
	teacherAttendence := &[]model.TeacherAttendence{}
	err := t.db.Model(&[]model.TeacherAttendence{}).Find(teacherAttendence).Error
	if err != nil {
		return nil, errors.New("can't gets teacher Attendence")
	}
	return teacherAttendence, nil
}
