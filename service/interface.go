package service

import (
	"context"

	"github.com/shahzaibaziz/account-manager/client"
	genModel "github.com/shahzaibaziz/account-manager/gen/models"
	domain "github.com/shahzaibaziz/account-manager/models"
)

// Manager defines the available functions for the given service implementation.
type Manager interface {
	DeleteUser(ctx context.Context, id string) error
	FindUser(ctx context.Context, id string) (*genModel.UserResponse, error)
	GetUsers(ctx context.Context, filter map[string]interface{}) ([]*genModel.UserResponse, error)
	SaveUser(ctx context.Context, user *domain.User) (*genModel.UserResponse, error)
	UpdateUser(ctx context.Context, user *domain.User) (*genModel.UserResponse, error)
}

type service struct {
	mgr client.Manager
}

// NewService return new service object
func NewService(manager client.Manager) Manager {
	return &service{mgr: manager}
}
