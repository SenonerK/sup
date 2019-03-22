package main

import (
	"time"

	"github.com/micro/go-log"
	"github.com/micro/go-web"

	"github.com/senonerk/sup/api/chat/handler"
)

func main() {
	service := web.NewService(
		web.Name("senonerk.sup.api.chat"),
		web.Version("0.0.1"),
		web.Handler(handler.New()),
		web.RegisterTTL(time.Second*40),
		web.RegisterInterval(time.Second*20),
	)

	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
