package client

import (
	openapiclient "github.com/go-openapi/runtime/client"
	client "github.com/shahzaibaziz/user_operations/gen/client"
	"github.com/spf13/viper"

	"github.com/shahzaibaziz/account-manager/config"
	genModel "github.com/shahzaibaziz/account-manager/gen/models"
	domain "github.com/shahzaibaziz/account-manager/models"
)

// Manager interface for user operation client
type Manager interface {
	DeleteUser(userID string) error
	GetUser(id string) (*genModel.UserResponse, error)
	GetAllUsers(filter map[string]interface{}) ([]*genModel.UserResponse, error)
	SaveUser(user *domain.User) (*genModel.UserResponse, error)
	UpdateUser(user *domain.User) (*genModel.UserResponse, error)
}

func newUserOperationsClient() *client.UserOperations {
	if viper.GetString(config.UserOperationsHost) == "" {
		panic("UserOperationsHost config not found")
	}

	userOps := client.New(openapiclient.New(
		viper.GetString(config.UserOperationsHost),
		client.DefaultBasePath,
		client.DefaultSchemes), nil)
	return userOps
}

// NewClientManager create new client manager
func NewClientManager() Manager {
	return &userClient{mgr: newUserOperationsClient()}
}
