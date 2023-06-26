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
func (r *complainResolver) ID(ctx context.Context, obj *model.Complain) (string, error) {
	return strconv.Itoa(int(obj.ID)), nil
}

// StudentID is the resolver for the StudentID field.
func (r *complainResolver) StudentID(ctx context.Context, obj *model.Complain) (string, error) {
	return strconv.Itoa(int(obj.StudentID)), nil
}

// AddComplain is the resolver for the addComplain field.
func (r *mutationResolver) AddComplain(ctx context.Context, complain string, studentID string) (*model.Complain, error) {
	a, _ := strconv.Atoi(studentID)
	service := service.NewComplainService()
	complainobj, err := service.CreateComplain(&model.Complain{
		Complain:  complain,
		StudentID: uint(a),
	})
	if err != nil {
		return nil, err
	}
	return complainobj, nil
}

// UpdateComplain is the resolver for the updateComplain field.
func (r *mutationResolver) UpdateComplain(ctx context.Context, id string, complain string, studentID string) (*model.Complain, error) {
	a, _ := strconv.Atoi(id)
	b, _ := strconv.Atoi(studentID)
	service := service.NewComplainService()
	complainobj, err := service.UpdateComplain(&model.Complain{
		ID:        uint(a),
		Complain:  complain,
		StudentID: uint(b),
	})
	if err != nil {
		return nil, err
	}
	return complainobj, nil
}

// GetComplain is the resolver for the getComplain field.
func (r *queryResolver) GetComplain(ctx context.Context, id string) (*model.Complain, error) {
	service := service.NewComplainService()
	complainobj, err := service.GetComplain(id)
	if err != nil {
		return nil, err
	}
	return complainobj, nil
}

// DeleteComplain is the resolver for the deleteComplain field.
func (r *queryResolver) DeleteComplain(ctx context.Context, id string) (bool, error) {
	service := service.NewComplainService()
	_, err := service.DeleteComplain(id)
	if err != nil {
		return false, nil
	}
	return true, nil
}

// Complain returns generated.ComplainResolver implementation.
func (r *Resolver) Complain() generated.ComplainResolver { return &complainResolver{r} }

type complainResolver struct{ *Resolver }
