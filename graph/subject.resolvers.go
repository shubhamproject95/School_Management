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

// AddSubject is the resolver for the addSubject field.
func (r *mutationResolver) AddSubject(ctx context.Context, subject string) (*model.Subject, error) {
	service := service.NewSubjectService()
	subjectobj, err := service.CreateSubject(&model.Subject{
		Subject: subject,
	})
	if err != nil {
		return nil, err
	}
	return subjectobj, nil
}

// UpdateSubject is the resolver for the updateSubject field.
func (r *mutationResolver) UpdateSubject(ctx context.Context, id string, subject string) (*model.Subject, error) {
	a, _ := strconv.Atoi(id)
	service := service.NewSubjectService()
	subjectobj, err := service.UpdateSubject(&model.Subject{
		ID:      uint(a),
		Subject: subject,
	})
	if err != nil {
		return nil, err
	}
	return subjectobj, nil
}

// GetSubject is the resolver for the getSubject field.
func (r *queryResolver) GetSubject(ctx context.Context, id string) (*model.Subject, error) {
	service := service.NewSubjectService()
	subjectobj, err := service.GetSubject(id)
	if err != nil {
		return nil, err
	}
	return subjectobj, nil
}

// DeleteSubject is the resolver for the deleteSubject field.
func (r *queryResolver) DeleteSubject(ctx context.Context, id string) (bool, error) {
	service := service.NewSubjectService()
	_, err := service.DeleteSubject(id)
	if err != nil {
		return false, nil
	}
	return true, nil
}

// ID is the resolver for the ID field.
func (r *subjectResolver) ID(ctx context.Context, obj *model.Subject) (string, error) {
	return strconv.Itoa(int(obj.ID)), nil
}

// Subject returns generated.SubjectResolver implementation.
func (r *Resolver) Subject() generated.SubjectResolver { return &subjectResolver{r} }

type subjectResolver struct{ *Resolver }
