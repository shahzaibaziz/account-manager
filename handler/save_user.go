package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	runtime "github.com/shahzaibaziz/account-manager"
	genModel "github.com/shahzaibaziz/account-manager/gen/models"
	"github.com/shahzaibaziz/account-manager/gen/restapi/operations/service"
)

// NewSaveUserHandler custom handler of save new User
func NewSaveUserHandler(rt *runtime.Runtime) service.SaveUserHandler {
	return &saveUser{rt: rt}
}

type saveUser struct {
	rt *runtime.Runtime
}

// Handle implementation of custom handler
func (r *saveUser) Handle(params service.SaveUserParams) middleware.Responder {
	ctx := context.Background()

	user, err := r.rt.Service().SaveUser(ctx, toUserModel(params.User))
	if err != nil {
		log(ctx).Errorf("failed to register new user", err)

		return service.NewSaveUserDefault(http.StatusInternalServerError).WithPayload(&genModel.Error{
			Code:    swag.String(fmt.Sprintf("%v", http.StatusInternalServerError)),
			Message: swag.String(err.Error()),
		})
	}

	log(ctx).Infof("got user %v", user)
	return service.NewSaveUserCreated().WithPayload(user)
}
