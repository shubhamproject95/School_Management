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
func (r *classResolver) ID(ctx context.Context, obj *model.Class) (string, error) {
	return strconv.Itoa(int(obj.ID)), nil
}

// AddClass is the resolver for the addClass field.
func (r *mutationResolver) AddClass(ctx context.Context, class string) (*model.Class, error) {
	service := service.NewClassService()
	classObj, err := service.CreateClass(&model.Class{
		Class: class,
	})
	if err != nil {
		return nil, err
	}
	return classObj, nil
}

// UpdateClass is the resolver for the updateClass field.
func (r *mutationResolver) UpdateClass(ctx context.Context, id string, class string) (bool, error) {
	a, _ := strconv.Atoi(id)
	service := service.NewClassService()
	_, err := service.UpdateClass(&model.Class{
		ID:    uint(a),
		Class: class,
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetClass is the resolver for the getClass field.
func (r *queryResolver) GetClass(ctx context.Context, id string) (*model.Class, error) {
	service := service.NewClassService()
	class, err := service.GetClass(id)
	if err != nil {
		return nil, err
	}
	return class, nil
}

// DeleteClass is the resolver for the deleteClass field.
func (r *queryResolver) DeleteClass(ctx context.Context, id string) (bool, error) {
	service := service.NewClassService()
	_, err := service.DeleteClass(id)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Class returns generated.ClassResolver implementation.
func (r *Resolver) Class() generated.ClassResolver { return &classResolver{r} }

type classResolver struct{ *Resolver }
