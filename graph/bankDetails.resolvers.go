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

// ID is the resolver for the ID field.
func (r *bankDetailResolver) ID(ctx context.Context, obj *model.BankDetail) (string, error) {
	return strconv.Itoa(int(obj.ID)), nil
}

// StaffID is the resolver for the StaffID field.
func (r *bankDetailResolver) StaffID(ctx context.Context, obj *model.BankDetail) (string, error) {
	return strconv.Itoa(int(obj.StaffID)), nil
}

// IsDefault is the resolver for the IsDefault field.
func (r *bankDetailResolver) IsDefault(ctx context.Context, obj *model.BankDetail) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// AddBankDetail is the resolver for the addBankDetail field.
func (r *mutationResolver) AddBankDetail(ctx context.Context, staffID string, name string, bank string, bankAccount string, ifsc string, branchCode string, isDefault string) (*model.BankDetail, error) {
	a, _ := strconv.Atoi(staffID)
	service := service.NewBankDetailervice()
	bankdetails, err := service.CreateBankDetail(&model.BankDetail{
		StaffID:     uint(a),
		Name:        name,
		Bank:        bank,
		BankAccount: bankAccount,
		IFSC:        ifsc,
		BranchCode:  branchCode,
		IsDefault:   true,
	})
	if err != nil {
		return nil, err
	}
	return bankdetails, nil
}

// DeleteBankDetail is the resolver for the deleteBankDetail field.
func (r *queryResolver) DeleteBankDetail(ctx context.Context, id string) (bool, error) {
	service := service.NewBankDetailervice()
	_, err := service.DeleteBankDetail(id)
	if err != nil {
		return false, nil
	}
	return true, nil
}

// BankDetail returns generated.BankDetailResolver implementation.
func (r *Resolver) BankDetail() generated.BankDetailResolver { return &bankDetailResolver{r} }

type bankDetailResolver struct{ *Resolver }
