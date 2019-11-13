package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"ihomegit/ihome/service/getSms/handler"

	getSms "ihomegit/ihome/service/getSms/proto/getSms"
	"github.com/micro/go-micro/registry/consul"
	"ihomegit/ihome/service/getSms/model"
)

func main() {
	//服务发现用consul
	consulReg := consul.NewRegistry()

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.getSms"),
		micro.Version("latest"),
		micro.Registry(consulReg),
		micro.Address(":8088"),
	)

	// Initialise service
	service.Init()
	model.InitRedis()

	// Register Handler
	getSms.RegisterGetSmsHandler(service.Server(), new(handler.GetSms))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
