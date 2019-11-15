package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"

	updateUser "ihomegit/ihome/service/updateUser/proto/updateUser"
	"github.com/micro/go-micro/registry/consul"
	"ihomegit/ihome/service/updateUser/handler"
	"ihomegit/ihome/model"
)

func main() {
	registry:=consul.NewRegistry()
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.updateUser"),
		micro.Version("latest"),
		micro.Address(":8090"),
		micro.Registry(registry),
	)

	// Initialise service
	service.Init()
	model.InitDb()
	// Register Handler
	updateUser.RegisterUpdateUserHandler(service.Server(), new(handler.GetUserInfo))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
