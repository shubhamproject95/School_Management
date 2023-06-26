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
func (r *loginResolver) ID(ctx context.Context, obj *model.Login) (string, error) {
	return strconv.Itoa(int(obj.ID)), nil
}

// UserTypeID is the resolver for the UserTypeID field.
func (r *loginResolver) UserTypeID(ctx context.Context, obj *model.Login) (int, error) {
	return int(obj.UserTypeID), nil
}

// RoleID is the resolver for the RoleID field.
func (r *loginResolver) RoleID(ctx context.Context, obj *model.Login) (int, error) {
	return int(obj.RoleID), nil
}

// AddUser is the resolver for the addUser field.
func (r *mutationResolver) AddUser(ctx context.Context, userName string, password string, userType string, rolesID int) (*model.User, error) {
	service := service.NewUserService()
	user, err := service.CreateUser(&model.User{
		UserName: userName,
		Password: password,
		UserType: userType,
		RolesID:  uint(rolesID),
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, userName string, password string, userType string, rolesID int) (*model.User, error) {
	a, _ := strconv.Atoi(id)
	service := service.NewUserService()
	user, err := service.UpdateUser(&model.User{
		ID:       uint(a),
		UserName: userName,
		Password: password,
		UserType: userType,
		RolesID:  uint(rolesID),
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateResetKet is the resolver for the updateResetKet field.
func (r *mutationResolver) UpdateResetKet(ctx context.Context, id int, userID int, resetKey string, isUsed bool, updatePasswordID int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

// GenerateToken is the resolver for the generateToken field.
func (r *mutationResolver) GenerateToken(ctx context.Context, loginID int) (string, error) {
	service := service.NewUserService()
	login, err := service.GenerateToken(uint(loginID))
	if err != nil {
		return " ", nil
	}
	return login, nil
}

// AddLogin is the resolver for the addLogin field.
func (r *mutationResolver) AddLogin(ctx context.Context, userName string, password string, userType string, userTypeID int, roleID int, token string) (*model.Login, error) {
	service := service.NewLoginService()
	login, err := service.CreateLogin(&model.Login{
		UserName:   userName,
		Password:   password,
		UserType:   userType,
		UserTypeID: uint(userTypeID),
		RoleID:     uint(roleID),
		Token:      token,
	})
	if err != nil {
		return nil, err
	}
	return login, nil
}

// DeleteUser is the resolver for the deleteUser field.
func (r *queryResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	service := service.NewUserService()
	_, err := service.DeleteUser(id)
	if err != nil {
		return false, nil
	}
	return true, nil
}

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context, id string) (*model.User, error) {
	service := service.NewUserService()
	user, err := service.GetUser(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UserLogin is the resolver for the userLogin field.
func (r *queryResolver) UserLogin(ctx context.Context, userName string, password string) (*model.Login, error) {
	service := service.NewUserService()
	user, err := service.UserLogin(userName, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserByModel is the resolver for the getUserByModel field.
func (r *queryResolver) GetUserByModel(ctx context.Context, userName string, password string, userType string, rolesID int) (*model.User, error) {
	service := service.NewUserService()
	user, err := service.GetUserByModel(&model.User{
		UserName: userName,
		Password: password,
		UserType: userType,
		RolesID:  uint(rolesID),
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetResetKey is the resolver for the getResetKey field.
func (r *queryResolver) GetResetKey(ctx context.Context, resetKey string) (*model.ResetPassword, error) {
	service := service.NewUserService()
	reset, err := service.GetResetKey(resetKey)
	if err != nil {
		return nil, err
	}
	return reset, nil
}

// Auth is the resolver for the auth field.
func (r *queryResolver) Auth(ctx context.Context, token string) (*model.Login, error) {
	panic(fmt.Errorf("not implemented"))
}

// ID is the resolver for the ID field.
func (r *resetPasswordResolver) ID(ctx context.Context, obj *model.ResetPassword) (string, error) {
	return strconv.Itoa(int(obj.ID)), nil
}

// UserID is the resolver for the UserID field.
func (r *resetPasswordResolver) UserID(ctx context.Context, obj *model.ResetPassword) (int, error) {
	return int(obj.UserID), nil
}

// UpdatePasswordID is the resolver for the UpdatePasswordID field.
func (r *resetPasswordResolver) UpdatePasswordID(ctx context.Context, obj *model.ResetPassword) (int, error) {
	return int(obj.UpdatePasswordID), nil
}

// ID is the resolver for the ID field.
func (r *userResolver) ID(ctx context.Context, obj *model.User) (string, error) {
	return strconv.Itoa(int(obj.ID)), nil
}

// RolesID is the resolver for the RolesID field.
func (r *userResolver) RolesID(ctx context.Context, obj *model.User) (int, error) {
	return int(obj.RolesID), nil
}

// Login returns generated.LoginResolver implementation.
func (r *Resolver) Login() generated.LoginResolver { return &loginResolver{r} }

// ResetPassword returns generated.ResetPasswordResolver implementation.
func (r *Resolver) ResetPassword() generated.ResetPasswordResolver { return &resetPasswordResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type loginResolver struct{ *Resolver }
type resetPasswordResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
