package service

import (
	"context"

	"github.com/pkg/errors"

	genModel "github.com/shahzaibaziz/account-manager/gen/models"
	domain "github.com/shahzaibaziz/account-manager/models"
)

// SaveUser save user object
func (s *service) SaveUser(ctx context.Context, user *domain.User) (*genModel.UserResponse, error) {
	usr, err := s.mgr.SaveUser(user)
	if err != nil {
		log(ctx).Errorf("failed to save user", err)
		return nil, errors.Wrap(err, "failed to save user")
	}

	return usr, nil
}

// FindUser find user object using id
func (s *service) FindUser(ctx context.Context, id string) (*genModel.UserResponse, error) {
	user, err := s.mgr.GetUser(id)
	if err != nil {
		log(ctx).Errorf("failed to get user", err)
		return nil, errors.Wrap(err, "failed to get user")
	}

	return user, nil
}

// GetUsers fetch all users using filter
func (s *service) GetUsers(ctx context.Context, filter map[string]interface{}) ([]*genModel.UserResponse, error) {
	users, err := s.mgr.GetAllUsers(filter)
	if err != nil {
		log(ctx).Errorf("failed to get all users", err)
		return nil, errors.Wrap(err, "failed to get all users")
	}

	return users, nil
}

// DeleteUser delete user
func (s *service) DeleteUser(ctx context.Context, id string) error {
	if err := s.mgr.DeleteUser(id); err != nil {
		log(ctx).Errorf("failed to delete user", err)
		return errors.Wrap(err, "failed to delete user")
	}

	return nil
}

// UpdateUser update user's name and address
func (s *service) UpdateUser(ctx context.Context, user *domain.User) (*genModel.UserResponse, error) {
	usr, err := s.mgr.UpdateUser(user)
	if err != nil {
		log(ctx).Errorf("failed to update user", err)
		return nil, errors.Wrap(err, "failed to update user")
	}

	return usr, nil
}
