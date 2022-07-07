package handler

import (
	genModel "github.com/shahzaibaziz/account-manager/gen/models"
	domain "github.com/shahzaibaziz/account-manager/models"
)

func toUserModel(user *genModel.User) *domain.User {
	return &domain.User{
		Name:    *user.Name,
		Address: *user.Address,
		Email:   user.Email.String(),
	}
}
