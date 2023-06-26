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

// AddRole is the resolver for the addRole field.
func (r *mutationResolver) AddRole(ctx context.Context, role string, status string) (*model.Role, error) {
	service := service.NewRoleervice()
	roleobj, err := service.CreateRole(&model.Role{
		Role:   role,
		Status: status,
	})
	if err != nil {
		return nil, err
	}
	return roleobj, nil
}

// UpdateRole is the resolver for the updateRole field.
func (r *mutationResolver) UpdateRole(ctx context.Context, id string, role string, status string) (*model.Role, error) {
	a, _ := strconv.Atoi(id)
	service := service.NewRoleervice()
	roleobj, err := service.UpdateRole(&model.Role{
		ID:     uint(a),
		Role:   role,
		Status: status,
	})
	if err != nil {
		return nil, err
	}
	return roleobj, nil
}

// GetRole is the resolver for the getRole field.
func (r *queryResolver) GetRole(ctx context.Context, id string) (*model.Role, error) {
	service := service.NewRoleervice()
	roleobj, err := service.GetRole(id)
	if err != nil {
		return nil, err
	}
	return roleobj, nil
}

// DeleteRole is the resolver for the deleteRole field.
func (r *queryResolver) DeleteRole(ctx context.Context, id string) (bool, error) {
	service := service.NewRoleervice()
	_, err := service.DeleteRole(id)
	if err != nil {
		return false, nil
	}
	return true, nil
}

// ID is the resolver for the ID field.
func (r *roleResolver) ID(ctx context.Context, obj *model.Role) (string, error) {
	return strconv.Itoa(int(obj.ID)), nil
}

// Role returns generated.RoleResolver implementation.
func (r *Resolver) Role() generated.RoleResolver { return &roleResolver{r} }

type roleResolver struct{ *Resolver }
