package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"School_gql/graph/generated"
	"School_gql/pkg/model"
	"School_gql/service"
	"context"
	"fmt"
	"strconv"
)

// AddTeacherAttendence is the resolver for the addTeacherAttendence field.
func (r *mutationResolver) AddTeacherAttendence(ctx context.Context, teacherAttendence string, teacherID string) (*model.TeacherAttendence, error) {
	a, _ := strconv.Atoi(teacherID)
	service := service.NewTeacherAttendence()
	teacher, err := service.CreateTeacherAttendence(&model.TeacherAttendence{
		TeacherAttendence: teacherAttendence,
		TeacherID:         uint(a),
	})
	if err != nil {
		return nil, err
	}
	return teacher, nil
}

// UpdateTeacherAttendence is the resolver for the updateTeacherAttendence field.
func (r *mutationResolver) UpdateTeacherAttendence(ctx context.Context, id string, teacherAttendence string, teacherID string) (*model.TeacherAttendence, error) {
	a, _ := strconv.Atoi(id)
	b, _ := strconv.Atoi(teacherID)
	service := service.NewTeacherAttendence()
	teacher, err := service.UpdateTeacherAttendence(&model.TeacherAttendence{
		ID:                uint(a),
		TeacherID:         uint(b),
		TeacherAttendence: teacherAttendence,
	})
	if err != nil {
		return nil, err
	}
	return teacher, nil
}

// GetTeacherAttendence is the resolver for the getTeacherAttendence field.
func (r *queryResolver) GetTeacherAttendence(ctx context.Context, id string) (*model.TeacherAttendence, error) {
	service := service.NewTeacherAttendence()
	teacher, err := service.GetTeacherAttendence(id)
	if err != nil {
		return nil, err
	}
	return teacher, nil
}

// DeleteTeacherAttendence is the resolver for the deleteTeacherAttendence field.
func (r *queryResolver) DeleteTeacherAttendence(ctx context.Context, id string) (bool, error) {
	service := service.NewTeacherAttendence()
	_, err := service.DeleteTeacherAttendence(id)
	if err != nil {
		return false, nil
	}
	return true, nil
}

// ID is the resolver for the ID field.
func (r *teacherAttendenceResolver) ID(ctx context.Context, obj *model.TeacherAttendence) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// TeacherID is the resolver for the TeacherID field.
func (r *teacherAttendenceResolver) TeacherID(ctx context.Context, obj *model.TeacherAttendence) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// TeacherAttendence returns generated.TeacherAttendenceResolver implementation.
func (r *Resolver) TeacherAttendence() generated.TeacherAttendenceResolver {
	return &teacherAttendenceResolver{r}
}

type teacherAttendenceResolver struct{ *Resolver }
