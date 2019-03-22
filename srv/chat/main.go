package main

import (
	"github.com/senonerk/sup/srv/auth/proto"
	"time"

	"github.com/senonerk/sup/srv/chat/models"

	"github.com/senonerk/sup/srv/chat/db"

	"github.com/micro/go-log"
	"github.com/micro/go-micro"

	"github.com/senonerk/sup/srv/chat/handler"
	proto "github.com/senonerk/sup/srv/chat/proto/chat"
)

func main() {
	service := micro.NewService(
		micro.Name("senonerk.sup.srv.chat"),
		micro.Version("0.0.1"),
		micro.RegisterTTL(time.Second*40),
		micro.RegisterInterval(time.Second*20),
	)

	// Connect to Database
	cnn, err := db.Connect()
	defer db.Close()
	if err != nil {
		log.Fatal(err)
		return
	}

	// Register models in ODM
	cnn.Register(&models.Chat{}, "chats")

	service.Init()

	proto.RegisterChatHandler(service.Server(), &handler.ChatService{
		Auth: auth.NewAuthService("senonerk.sup.srv.auth", service.Client()),
	})

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
