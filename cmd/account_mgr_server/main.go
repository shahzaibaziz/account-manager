package main

import (
	"github.com/go-openapi/loads"
	"github.com/sirupsen/logrus"

	runtime "github.com/shahzaibaziz/account-manager"
	"github.com/shahzaibaziz/account-manager/gen/restapi"
	"github.com/shahzaibaziz/account-manager/gen/restapi/operations"
	"github.com/shahzaibaziz/account-manager/handler"
)

func main() {
	log := logrus.WithField("pkg", "main")

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}
	rt, err := runtime.NewRuntime()
	if err != nil {
		log.Fatalln(err)
	}

	var server *restapi.Server

	api := operations.NewAccountManagerAPI(swaggerSpec)
	api.Logger = log.Infof

	handler.NewCustomHandler(api, rt)

	server = restapi.NewServer(api)
	server.Port = 9090

	defer server.Shutdown()

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
