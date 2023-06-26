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
func (r *feesResolver) ID(ctx context.Context, obj *model.Fees) (string, error) {
	return strconv.Itoa(int(obj.ID)), nil
}

// StudentID is the resolver for the StudentID field.
func (r *feesResolver) StudentID(ctx context.Context, obj *model.Fees) (int, error) {
	return int(obj.StudentID), nil
}

// FeesPaid is the resolver for the FeesPaid field.
func (r *feesResolver) FeesPaid(ctx context.Context, obj *model.Fees) (int, error) {
	return int(obj.FeesPaid), nil
}

// TotalFees is the resolver for the TotalFees field.
func (r *feesResolver) TotalFees(ctx context.Context, obj *model.Fees) (int, error) {
	return int(obj.TotalFees), nil
}

// Pending is the resolver for the Pending field.
func (r *feesResolver) Pending(ctx context.Context, obj *model.Fees) (int, error) {
	return int(obj.Pending), nil
}

// ClassID is the resolver for the ClassID field.
func (r *feesResolver) ClassID(ctx context.Context, obj *model.Fees) (int, error) {
	return int(obj.ClassID), nil
}

// AddFee is the resolver for the addFee field.
func (r *mutationResolver) AddFee(ctx context.Context, studentID int, session string, feesPaid int, totalFees int, pending int, month string, classID int) (*model.Fees, error) {
	service := service.NewFeeService()
	fee, err := service.CreateFees(&model.Fees{
		StudentID: uint(studentID),
		Session:   session,
		FeesPaid:  uint(feesPaid),
		TotalFees: uint(totalFees),
		Pending:   uint(pending),
		Month:     month,
		ClassID:   uint(classID),
	})
	if err != nil {
		return nil, err
	}
	return fee, nil
}

// UpdateFee is the resolver for the updateFee field.
func (r *mutationResolver) UpdateFee(ctx context.Context, id string, studentID *int, session *string, feesPaid *int, totalFees *int, pending *int, month *string, classID *int) (bool, error) {
	a, _ := strconv.Atoi(id)
	service := service.NewFeeService()
	_, err := service.UpdateFees(&model.Fees{
		ID:        uint(a),
		StudentID: uint(*studentID),
		Session:   *session,
		FeesPaid:  uint(*feesPaid),
		TotalFees: uint(*totalFees),
		Pending:   uint(*pending),
		Month:     *month,
		ClassID:   uint(*classID),
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetFee is the resolver for the getFee field.
func (r *queryResolver) GetFee(ctx context.Context, id string) (*model.Fees, error) {
	service := service.NewFeeService()
	fee, err := service.GetFees(id)
	if err != nil {
		return nil, err
	}
	return fee, nil
}

// DeleteFee is the resolver for the deleteFee field.
func (r *queryResolver) DeleteFee(ctx context.Context, id string) (bool, error) {
	service := service.NewFeeService()
	_, err := service.DeleteFees(id)
	if err != nil {
		return false, nil
	}
	return true, nil
}

// Fees returns generated.FeesResolver implementation.
func (r *Resolver) Fees() generated.FeesResolver { return &feesResolver{r} }

type feesResolver struct{ *Resolver }
