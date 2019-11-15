package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"ihomegit/ihome/service/RegAndLog/handler"

	RegAndLog "ihomegit/ihome/service/RegAndLog/proto/RegAndLog"
	"ihomegit/ihome/service/RegAndLog/model"
	"github.com/micro/go-micro/registry/consul"
)

func main() {
	consulReg :=consul.NewRegistry();

	model.InitRedis();
    model.InitDb();
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.RegAndLog"),
		micro.Version("latest"),
		micro.Registry(consulReg),
		micro.Address(":8089"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	RegAndLog.RegisterRegAndLogHandler(service.Server(), new(handler.Register))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
