package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/senonerk/sup/srv/auth/handler"

	proto "github.com/senonerk/sup/srv/auth/proto"
)

func main() {
	service := micro.NewService(
		micro.Name("senonerk.sup.srv.auth"),
		micro.Version("0.0.1"),
	)

	service.Init()

	proto.RegisterAuthHandler(service.Server(), handler.New())

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
