package main

import (
	"time"

	"github.com/senonerk/sup/srv/auth/db"

	"github.com/micro/go-log"
	"github.com/micro/go-micro"

	"github.com/senonerk/sup/srv/auth/handler"
	"github.com/senonerk/sup/srv/auth/models"
	proto "github.com/senonerk/sup/srv/auth/proto"
)

func main() {
	service := micro.NewService(
		micro.Name("senonerk.sup.srv.auth"),
		micro.Version("0.0.1"),
		micro.RegisterTTL(time.Second*40),
		micro.RegisterInterval(time.Second*20),
	)

	NewUserPub := micro.NewPublisher("sup.auth.NewUser", service.Client())

	// Connect to Database
	cnn, err := db.Connect()
	defer db.Close()
	if err != nil {
		log.Fatal(err)
		return
	}

	// Register models in ODM
	cnn.Register(&models.User{}, "users")

	//--- SEED DB ---//
	// m := cnn.Model("User")
	// usr := &models.User{
	// 	UserName: "test",
	// 	Password: "fo",
	// 	Permissions: []models.Permission{
	// 		models.Permission{
	// 			Tag:   "ALL",
	// 			Grant: true,
	// 		},
	// 	},
	// }
	// m.New(usr)
	// usr.Save()

	service.Init()

	proto.RegisterAuthHandler(service.Server(), handler.New(NewUserPub))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
