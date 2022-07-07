package runtime

import (
	domainClient "github.com/shahzaibaziz/account-manager/client"
	"github.com/shahzaibaziz/account-manager/service"
)

// Runtime initializes values for entry point to our application
type Runtime struct {
	service service.Manager
}

// NewRuntime creates a new runtime
func NewRuntime() (*Runtime, error) {
	return &Runtime{service: service.NewService(domainClient.NewClientManager())}, nil
}

// Service return  service layer object
func (rt *Runtime) Service() service.Manager {
	return rt.service
}
