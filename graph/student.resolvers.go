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

// AddStudent is the resolver for the addStudent field.
func (r *mutationResolver) AddStudent(ctx context.Context, name string, address string, dob string, fatherName string, motherName string, status string) (*model.Student, error) {
	service := service.NewStudentService()
	student, err := service.CreateStudent(&model.Student{
		Name:        name,
		Address:     address,
		DOB:         dob,
		Father_Name: fatherName,
		Mother_Name: motherName,
		Status:      status,
	})
	if err != nil {
		return nil, err
	}
	return student, nil
}

// UpdateStudent is the resolver for the updateStudent field.
func (r *mutationResolver) UpdateStudent(ctx context.Context, id string, name string, address string, dob string, fatherName string, motherName string, status string) (bool, error) {
	a, _ := strconv.Atoi(id)
	service := service.NewStudentService()
	_, err := service.UpdateStudent(&model.Student{
		ID:          uint(a),
		Name:        name,
		Address:     address,
		DOB:         dob,
		Father_Name: fatherName,
		Mother_Name: motherName,
		Status:      status,
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetStudent is the resolver for the getStudent field.
func (r *queryResolver) GetStudent(ctx context.Context, id string) (*model.Student, error) {
	service := service.NewStudentService()

	student, err := service.GetStudent(id)
	if err != nil {
		return nil, err
	}
	return student, nil
}

// DeleteStudent is the resolver for the deleteStudent field.
func (r *queryResolver) DeleteStudent(ctx context.Context, id string) (bool, error) {
	service := service.NewStudentService()
	_, err := service.DeleteStudent(id)
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetStudents is the resolver for the getStudents field.
func (r *queryResolver) GetStudents(ctx context.Context) ([]*model.Student, error) {
	// students := []*model.Student{}
	// service := service.NewStudentService()
	// students, err := service.GetsStudent()
	// if err != nil {
	// 	return nil, err
	// }
	return nil, nil
}

// ID is the resolver for the ID field.
func (r *studentResolver) ID(ctx context.Context, obj *model.Student) (string, error) {
	return strconv.Itoa(int(obj.ID)), nil
}

// Student returns generated.StudentResolver implementation.
func (r *Resolver) Student() generated.StudentResolver { return &studentResolver{r} }

type studentResolver struct{ *Resolver }
