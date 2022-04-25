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

// NewGetUserHandler custom handler for get  user
func NewGetUserHandler(rt *runtime.Runtime) service.GetUserHandler {
	return &getUser{rt: rt}
}

type getUser struct {
	rt *runtime.Runtime
}

// Handle implementation of custom handler
func (r *getUser) Handle(params service.GetUserParams) middleware.Responder {
	ctx := context.Background()

	user, err := r.rt.Service().FindUser(ctx, params.UserID)
	if err != nil {
		log(ctx).Errorf("failed to get user", err)

		return service.NewGetUserDefault(http.StatusInternalServerError).WithPayload(
			&genModel.Error{
				Code:    swag.String(fmt.Sprintf("%v", http.StatusInternalServerError)),
				Message: swag.String(err.Error()),
			})
	}

	log(ctx).Infof("user found %v", user)
	return service.NewGetUserOK().WithPayload(user)
}
