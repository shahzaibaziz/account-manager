package client

import (
	"github.com/shahzaibaziz/user_operations/gen/client"
	params "github.com/shahzaibaziz/user_operations/gen/client/service"

	genModel "github.com/shahzaibaziz/account-manager/gen/models"
	domain "github.com/shahzaibaziz/account-manager/models"
)

type userClient struct {
	mgr *client.UserOperations
}

// DeleteUser delete user from user operation server
func (c *userClient) DeleteUser(userID string) error {
	_, err := c.mgr.Service.DeleteUser(params.NewDeleteUserParams().WithUserID(userID))
	if err != nil {
		log().Errorf("failed to delete user ", err)
		return err
	}

	return nil
}

// GetUser get user from user operations service
func (c *userClient) GetUser(id string) (*genModel.UserResponse, error) {
	res, err := c.mgr.Service.GetUser(params.NewGetUserParams().WithUserID(id))
	if err != nil {
		return nil, err
	}

	return asUserResponse(res.Payload), nil
}

// GetAllUsers get all users from user operations base on given filter
func (c *userClient) GetAllUsers(filter map[string]interface{}) ([]*genModel.UserResponse, error) {
	params := params.NewGetAllUsersParams()

	// update filter with the correct field name and value.
	if val, ok := filter["name"].(*string); ok {
		params.WithName(val)
	}
	if val, ok := filter["limit"].(*int64); ok {
		params.WithLimit(val)
	}

	res, _, err := c.mgr.Service.GetAllUsers(params)
	if err != nil {
		return nil, err
	}

	users := make([]*genModel.UserResponse, 0)
	for _, user := range res.Payload {
		users = append(users, asUserResponse(user))
	}

	return users, nil
}

// SaveUser save new user in user operation service
func (c *userClient) SaveUser(user *domain.User) (*genModel.UserResponse, error) {
	res, err := c.mgr.Service.RegisterUser(params.NewRegisterUserParams().WithUser(asOperationUser(user)))
	if err != nil {
		return nil, err
	}

	return asUserResponse(res.Payload), nil
}

// UpdateUser update user in user operation service
func (c *userClient) UpdateUser(user *domain.User) (*genModel.UserResponse, error) {
	apiUser := asOperationUpdateUser(user)
	res, err := c.mgr.Service.UpdateUser(params.NewUpdateUserParams().WithUser(apiUser).WithUserID(user.ID))
	if err != nil {
		return nil, err
	}

	return asUserResponse(res.Payload), nil
}
