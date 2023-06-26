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

// AddStudentAttendence is the resolver for the addStudentAttendence field.
func (r *mutationResolver) AddStudentAttendence(ctx context.Context, attendence string, studentID string) (*model.StudentAttendence, error) {
	a, _ := strconv.Atoi(studentID)
	service := service.NewAttendenceService()
	attendenceobj, err := service.CreateAttendence(&model.StudentAttendence{
		Attendence: attendence,
		StudentID:  uint(a),
	})
	if err != nil {
		return nil, err
	}
	return attendenceobj, nil
}

// UpdateStudentAttendence is the resolver for the updateStudentAttendence field.
func (r *mutationResolver) UpdateStudentAttendence(ctx context.Context, id string, attendence string, studentID string) (*model.StudentAttendence, error) {
	a, _ := strconv.Atoi(id)
	b, _ := strconv.Atoi(studentID)
	service := service.NewAttendenceService()
	attendenceobj, err := service.UpdateAttendence(&model.StudentAttendence{
		ID:         uint(a),
		Attendence: attendence,
		StudentID:  uint(b),
	})
	if err != nil {
		return nil, err
	}
	return attendenceobj, nil
}

// GetStudentAttendence is the resolver for the getStudentAttendence field.
func (r *queryResolver) GetStudentAttendence(ctx context.Context, id string) (*model.StudentAttendence, error) {
	service := service.NewAttendenceService()
	attendenceobj, err := service.GetAttendence(id)
	if err != nil {
		return nil, err
	}
	return attendenceobj, nil
}

// DeleteStudentAttendence is the resolver for the deleteStudentAttendence field.
func (r *queryResolver) DeleteStudentAttendence(ctx context.Context, id string) (bool, error) {
	service := service.NewAttendenceService()
	_, err := service.DeleteAttendence(id)
	if err != nil {
		return false, nil
	}
	return true, nil
}

// ID is the resolver for the ID field.
func (r *studentAttendenceResolver) ID(ctx context.Context, obj *model.StudentAttendence) (string, error) {
	return strconv.Itoa(int(obj.ID)), nil
}

// StudentID is the resolver for the StudentID field.
func (r *studentAttendenceResolver) StudentID(ctx context.Context, obj *model.StudentAttendence) (string, error) {
	return strconv.Itoa(int(obj.ID)), nil
}

// StudentAttendence returns generated.StudentAttendenceResolver implementation.
func (r *Resolver) StudentAttendence() generated.StudentAttendenceResolver {
	return &studentAttendenceResolver{r}
}

type studentAttendenceResolver struct{ *Resolver }
