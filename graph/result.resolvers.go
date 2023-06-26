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

// AddResult is the resolver for the addResult field.
func (r *mutationResolver) AddResult(ctx context.Context, resultType string, rollNumber string) (*model.Result, error) {
	a, _ := strconv.Atoi(rollNumber)
	service := service.NewResultService()
	result, err := service.CreateResult(&model.Result{
		ResultType: resultType,
		RollNumber: a,
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateResult is the resolver for the updateResult field.
func (r *mutationResolver) UpdateResult(ctx context.Context, resultType string, rollNumber string) (*model.Result, error) {
	a, _ := strconv.Atoi(rollNumber)
	service := service.NewResultService()
	result, err := service.UpdateResult(&model.Result{
		ResultType: resultType,
		RollNumber: a,
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetResult is the resolver for the getResult field.
func (r *queryResolver) GetResult(ctx context.Context, id string) (*model.Result, error) {
	service := service.NewResultService()
	result, err := service.GetResult(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteResult is the resolver for the deleteResult field.
func (r *queryResolver) DeleteResult(ctx context.Context, id string) (bool, error) {
	service := service.NewResultService()
	_, err := service.DeleteResult(id)
	if err != nil {
		return false, nil
	}
	return true, nil
}

// ID is the resolver for the ID field.
func (r *resultResolver) ID(ctx context.Context, obj *model.Result) (string, error) {
	return strconv.Itoa(int(obj.ID)), nil
}

// Result returns generated.ResultResolver implementation.
func (r *Resolver) Result() generated.ResultResolver { return &resultResolver{r} }

type resultResolver struct{ *Resolver }
