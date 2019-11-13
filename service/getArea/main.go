package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"ihomegit/ihome/service/getArea/handler"

	getArea "ihomegit/ihome/service/getArea/proto/getArea"
	"ihomegit/ihome/service/getArea/model"
)

func main() {
	model.InitDb()
	model.InitRedis()
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.getArea"),
		micro.Version("latest"),
		micro.Address(":8086"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	getArea.RegisterGetAreaHandler(service.Server(), new(handler.GetArea))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
