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
func (r *homeworkResolver) ID(ctx context.Context, obj *model.Homework) (string, error) {
	return strconv.Itoa(int(obj.ID)), nil
}

// SubjectID is the resolver for the SubjectID field.
func (r *homeworkResolver) SubjectID(ctx context.Context, obj *model.Homework) (int, error) {
	return int(obj.SubjectID), nil
}

// TeacherID is the resolver for the TeacherID field.
func (r *homeworkResolver) TeacherID(ctx context.Context, obj *model.Homework) (int, error) {
	return int(obj.TeacherID), nil
}

// ClassID is the resolver for the ClassID field.
func (r *homeworkResolver) ClassID(ctx context.Context, obj *model.Homework) (int, error) {
	return int(obj.ClassID), nil
}

// AddHomework is the resolver for the addHomework field.
func (r *mutationResolver) AddHomework(ctx context.Context, subjectID int, teacherID int, classID int, description string, submissionDate string) (*model.Homework, error) {
	service := service.NewHomeworkService()
	homework, err := service.CreateHomework(&model.Homework{
		SubjectID:      uint(subjectID),
		ClassID:        uint(classID),
		TeacherID:      uint(teacherID),
		Description:    description,
		SubmissionDate: submissionDate,
	})
	if err != nil {
		return nil, err
	}
	return homework, nil
}

// DeleteHomework is the resolver for the deleteHomework field.
func (r *queryResolver) DeleteHomework(ctx context.Context, id string) (bool, error) {
	service := service.NewHomeworkService()
	_, err := service.DeleteHomework(id)
	if err != nil {
		return false, nil
	}
	return true, nil
}

// Homework returns generated.HomeworkResolver implementation.
func (r *Resolver) Homework() generated.HomeworkResolver { return &homeworkResolver{r} }

type homeworkResolver struct{ *Resolver }
