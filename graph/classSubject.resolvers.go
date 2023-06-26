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
func (r *classSubjectResolver) ID(ctx context.Context, obj *model.ClassSubject) (string, error) {
	return strconv.Itoa(int(obj.ID)), nil
}

// ClassID is the resolver for the ClassID field.
func (r *classSubjectResolver) ClassID(ctx context.Context, obj *model.ClassSubject) (int, error) {
	return int(obj.ClassID), nil
}

// SubjectID is the resolver for the SubjectID field.
func (r *classSubjectResolver) SubjectID(ctx context.Context, obj *model.ClassSubject) (int, error) {
	return int(obj.SubjectID), nil
}

// AddClassSubject is the resolver for the addClassSubject field.
func (r *mutationResolver) AddClassSubject(ctx context.Context, classID int, subjectID int) (*model.ClassSubject, error) {
	service := service.NewClassSubjectService()
	class, err := service.CreateClassSubject(&model.ClassSubject{
		ClassID:   uint(classID),
		SubjectID: uint(subjectID),
	})
	if err != nil {
		return nil, err
	}
	return class, nil
}

// UpdateClassSubject is the resolver for the updateClassSubject field.
func (r *mutationResolver) UpdateClassSubject(ctx context.Context, id string, classID int, subjectID int) (bool, error) {
	a, _ := strconv.Atoi(id)
	service := service.NewClassSubjectService()
	_, err := service.UpdateClassSubject(&model.ClassSubject{
		ID:        uint(a),
		ClassID:   uint(classID),
		SubjectID: uint(subjectID),
	})
	if err != nil {
		return false, nil
	}
	return true, nil
}

// GetClassSubject is the resolver for the getClassSubject field.
func (r *queryResolver) GetClassSubject(ctx context.Context, id string) (*model.ClassSubject, error) {
	a, _ := strconv.Atoi(id)
	service := service.NewClassSubjectService()
	class, err := service.GetClassSubject(int(a))
	if err != nil {
		return nil, err
	}
	return class, nil
}

// DeleteClassSubject is the resolver for the deleteClassSubject field.
func (r *queryResolver) DeleteClassSubject(ctx context.Context, id string) (bool, error) {
	a, _ := strconv.Atoi(id)
	service := service.NewClassSubjectService()
	_, err := service.DeleteClassSubject(int(a))
	if err != nil {
		return false, nil
	}
	return true, nil
}

// ClassSubject returns generated.ClassSubjectResolver implementation.
func (r *Resolver) ClassSubject() generated.ClassSubjectResolver { return &classSubjectResolver{r} }

type classSubjectResolver struct{ *Resolver }
