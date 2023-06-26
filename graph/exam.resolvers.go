package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"School_gql/graph/generated"
	"School_gql/pkg/model"
	"School_gql/service"
	"context"
	"strconv"
)

// ID is the resolver for the ID field.
func (r *examResolver) ID(ctx context.Context, obj *model.Exam) (string, error) {
	return strconv.Itoa(int(obj.ID)), nil
}

// SubjectID is the resolver for the SubjectID field.
func (r *examResolver) SubjectID(ctx context.Context, obj *model.Exam) (string, error) {
	return strconv.Itoa(int(obj.SubjectID)), nil
}

// ClassID is the resolver for the ClassID field.
func (r *examResolver) ClassID(ctx context.Context, obj *model.Exam) (string, error) {
	return strconv.Itoa(int(obj.ClassID)), nil
}

// InternalMarks is the resolver for the InternalMarks field.
func (r *examResolver) InternalMarks(ctx context.Context, obj *model.Exam) (int, error) {
	return int(obj.InternalMarks), nil
}

// ExtenalMarks is the resolver for the ExtenalMarks field.
func (r *examResolver) ExtenalMarks(ctx context.Context, obj *model.Exam) (int, error) {
	return int(obj.ExtenalMarks), nil
}

// PracticalMarks is the resolver for the PracticalMarks field.
func (r *examResolver) PracticalMarks(ctx context.Context, obj *model.Exam) (int, error) {
	return int(obj.PracticalMarks), nil
}

// AddExam is the resolver for the addExam field.
func (r *mutationResolver) AddExam(ctx context.Context, subjectID string, classID string, examType string, internalMarks int, extenalMarks int, practicalMarks int) (*model.Exam, error) {
	a, _ := strconv.Atoi(subjectID)
	b, _ := strconv.Atoi(classID)

	service := service.NewExamService()
	exam, err := service.CreateExam(&model.Exam{
		SubjectID:      uint(a),
		ClassID:        uint(b),
		ExamType:       examType,
		InternalMarks:  uint(internalMarks),
		ExtenalMarks:   uint(extenalMarks),
		PracticalMarks: uint(practicalMarks),
	})
	if err != nil {
		return nil, err
	}
	return exam, nil
}

// UpdateExam is the resolver for the updateExam field.
func (r *mutationResolver) UpdateExam(ctx context.Context, id string, subjectID string, classID string, examType string, internalMarks int, extenalMarks int, practicalMarks int) (*model.Exam, error) {
	a, _ := strconv.Atoi(subjectID)
	b, _ := strconv.Atoi(classID)
	f, _ := strconv.Atoi(id)

	service := service.NewExamService()
	exam, err := service.UpdateExam(&model.Exam{
		ID:             uint(f),
		SubjectID:      uint(a),
		ClassID:        uint(b),
		ExamType:       examType,
		InternalMarks:  uint(internalMarks),
		ExtenalMarks:   uint(extenalMarks),
		PracticalMarks: uint(practicalMarks),
	})
	if err != nil {
		return nil, err
	}
	return exam, nil
}

// GetExam is the resolver for the getExam field.
func (r *queryResolver) GetExam(ctx context.Context, id string) (*model.Exam, error) {
	service := service.NewExamService()
	exam, err := service.GetExam(id)
	if err != nil {
		return nil, err
	}
	return exam, nil
}

// DeleteExam is the resolver for the deleteExam field.
func (r *queryResolver) DeleteExam(ctx context.Context, id string) (bool, error) {
	service := service.NewExamService()
	_, err := service.DeleteExam(id)
	if err != nil {
		return false, nil
	}
	return true, nil
}

// Exam returns generated.ExamResolver implementation.
func (r *Resolver) Exam() generated.ExamResolver { return &examResolver{r} }

type examResolver struct{ *Resolver }
