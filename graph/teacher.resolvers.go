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

// AddTeacher is the resolver for the addTeacher field.
func (r *mutationResolver) AddTeacher(ctx context.Context, firstName string, lastName string, department string, dob string, joiningAt string, status string) (*model.Teacher, error) {
	service := service.NewTeacherService()
	teacher, err := service.CreateTeacher(&model.Teacher{
		FirstName:  firstName,
		LastName:   lastName,
		Department: department,
		DOB:        dob,
		JoiningAt:  joiningAt,
		Status:     status,
	})
	if err != nil {
		return nil, err
	}
	return teacher, nil
}

// UpdateTeacher is the resolver for the updateTeacher field.
func (r *mutationResolver) UpdateTeacher(ctx context.Context, id string, firstName string, lastName string, department string, dob string, joiningAt string, status string) (*model.Teacher, error) {
	a, _ := strconv.Atoi(id)
	service := service.NewTeacherService()
	teacher, err := service.UpdateTeacher(&model.Teacher{
		ID:         uint(a),
		FirstName:  firstName,
		LastName:   lastName,
		Department: department,
		DOB:        dob,
		JoiningAt:  joiningAt,
		Status:     status,
	})
	if err != nil {
		return nil, err
	}
	return teacher, nil
}

// GetTeacher is the resolver for the getTeacher field.
func (r *queryResolver) GetTeacher(ctx context.Context, id string) (*model.Teacher, error) {
	service := service.NewTeacherService()
	teacher, err := service.GetTeacher(id)
	if err != nil {
		return nil, err
	}
	return teacher, nil
}

// DeleteTeacher is the resolver for the deleteTeacher field.
func (r *queryResolver) DeleteTeacher(ctx context.Context, id string) (bool, error) {
	service := service.NewTeacherService()
	_, err := service.DeleteTeacher(id)
	if err != nil {
		return false, nil
	}
	return true, nil
}

// ID is the resolver for the ID field.
func (r *teacherResolver) ID(ctx context.Context, obj *model.Teacher) (string, error) {
	return strconv.Itoa(int(obj.ID)), nil
}

// Teacher returns generated.TeacherResolver implementation.
func (r *Resolver) Teacher() generated.TeacherResolver { return &teacherResolver{r} }

type teacherResolver struct{ *Resolver }
