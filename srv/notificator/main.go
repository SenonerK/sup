package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/senonerk/sup/srv/notificator/handler"

	"github.com/senonerk/sup/srv/notificator/proto/notificator"
)

func main() {
	service := micro.NewService(
		micro.Name("senonerk.sup.srv.notificator"),
		micro.Version("0.0.1"),
	)

	service.Init()

	notify.RegisterNotificatorHandler(service.Server(), new(handler.NotifyService))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
