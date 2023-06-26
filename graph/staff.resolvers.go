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

// AddStaff is the resolver for the addStaff field.
func (r *mutationResolver) AddStaff(ctx context.Context, name string, dob string, joiningDate string, aadharno string, staffType string) (*model.Staff, error) {
	service := service.NewStaffService()
	staff, err := service.CreateStaff(&model.Staff{
		Name:        name,
		DOB:         dob,
		JoiningDate: joiningDate,
		Aadharno:    aadharno,
		StaffType:   staffType,
	})
	if err != nil {
		return nil, err
	}
	return staff, nil
}

// UpdateStaff is the resolver for the updateStaff field.
func (r *mutationResolver) UpdateStaff(ctx context.Context, id string, name string, dob string, joiningDate string, aadharno string, staffType string) (*model.Staff, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetStaff is the resolver for the getStaff field.
func (r *queryResolver) GetStaff(ctx context.Context, staffType string) (*model.Staff, error) {
	panic(fmt.Errorf("not implemented"))
}

// DeleteStaff is the resolver for the deleteStaff field.
func (r *queryResolver) DeleteStaff(ctx context.Context, id string) (bool, error) {
	service := service.NewStaffService()
	_, err := service.DeleteStaff(id)
	if err != nil {
		return false, nil
	}
	return true, nil
}

// GetStaffByID is the resolver for the getStaffByID field.
func (r *queryResolver) GetStaffByID(ctx context.Context, id string, staffType string) (*model.Staff, error) {
	service := service.NewStaffService()
	staff, err := service.GetStaffByID("13", staffType)
	if err != nil {
		return nil, err
	}
	return staff, nil
}

// ID is the resolver for the ID field.
func (r *staffResolver) ID(ctx context.Context, obj *model.Staff) (string, error) {
	return strconv.Itoa(int(obj.ID)), nil
}

// Staff returns generated.StaffResolver implementation.
func (r *Resolver) Staff() generated.StaffResolver { return &staffResolver{r} }

type staffResolver struct{ *Resolver }
