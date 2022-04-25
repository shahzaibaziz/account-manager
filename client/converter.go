package client

import (
	"github.com/go-openapi/strfmt"
	genModel "github.com/shahzaibaziz/account-manager/gen/models"
	domain "github.com/shahzaibaziz/account-manager/models"
	userOps "github.com/shahzaibaziz/user_operations/gen/models"
)

func asUserResponse(user *userOps.UserResponse) *genModel.UserResponse {
	return &genModel.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Address:   user.Address,
		CreatedAt: user.CreatedAt,
	}
}

func asOperationUser(user *domain.User) *userOps.User {
	email := strfmt.Email(user.Email)
	return &userOps.User{
		Name:    &user.Name,
		Email:   &email,
		Address: &user.Address,
	}
}

func asOperationUpdateUser(user *domain.User) *userOps.UpdateUser {
	return &userOps.UpdateUser{
		Name:    user.Name,
		Address: user.Address,
	}
}
