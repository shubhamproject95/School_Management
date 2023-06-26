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
func (r *guardianResolver) ID(ctx context.Context, obj *model.Guardian) (string, error) {
	return strconv.Itoa(int(obj.ID)), nil
}

// StudentID is the resolver for the StudentID field.
func (r *guardianResolver) StudentID(ctx context.Context, obj *model.Guardian) (int, error) {
	return int(obj.StudentID), nil
}

// AddGuardian is the resolver for the addGuardian field.
func (r *mutationResolver) AddGuardian(ctx context.Context, studentID int, guardianType string, guardianName string) (*model.Guardian, error) {
	service := service.NewGuardianService()
	guardian, err := service.CreateGuardian(&model.Guardian{
		StudentID:    uint(studentID),
		GuardianType: guardianType,
		GuardianName: guardianName,
	})
	if err != nil {
		return nil, err
	}
	return guardian, nil
}

// DeleteGuardian is the resolver for the deleteGuardian field.
func (r *queryResolver) DeleteGuardian(ctx context.Context, id string) (bool, error) {
	service := service.NewGuardianService()
	_, err := service.DeleteGuardian(id)
	if err != nil {
		return false, nil
	}
	return true, nil
}

// Guardian returns generated.GuardianResolver implementation.
func (r *Resolver) Guardian() generated.GuardianResolver { return &guardianResolver{r} }

type guardianResolver struct{ *Resolver }
