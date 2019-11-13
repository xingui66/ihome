package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"ihomegit/ihome/service/getImg/handler"

	getImg "ihomegit/ihome/service/getImg/proto/getImg"
	"ihomegit/ihome/service/getImg/model"
)

func main() {
	model.InitRedis()
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.getImg"),
		micro.Version("latest"),
		micro.Address(":8087"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	getImg.RegisterGetImgHandler(service.Server(), new(handler.GetImg))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
