package handler

import (
	runtime "github.com/shahzaibaziz/account-manager"
	"github.com/shahzaibaziz/account-manager/gen/restapi/operations"
)

// NewCustomHandler mapping custom handler to auto generated handler
func NewCustomHandler(api *operations.AccountManagerAPI, rt *runtime.Runtime) {
	// custom handler
	api.ServiceSaveUserHandler = NewSaveUserHandler(rt)
	api.ServiceGetUserHandler = NewGetUserHandler(rt)
	api.ServiceUpdateUserHandler = NewUpdateUserHandler(rt)
	api.ServiceGetAllUsersHandler = NewGetAllUsers(rt)
	api.ServiceDeleteUserHandler = NewDeleteUserHandler(rt)
}
