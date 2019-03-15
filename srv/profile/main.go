package main

import (
	"time"

	"github.com/senonerk/sup/srv/auth/proto"

	"github.com/senonerk/sup/srv/profile/models"

	"github.com/senonerk/sup/srv/profile/db"

	"github.com/senonerk/sup/srv/profile/proto/profile"

	"github.com/senonerk/sup/srv/profile/handler"

	"github.com/micro/go-log"
	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
		micro.Name("senonerk.sup.srv.profile"),
		micro.Version("0.0.1"),
		micro.RegisterTTL(time.Second*40),
		micro.RegisterInterval(time.Second*20),
	)

	if err := db.Connect(); err != nil {
		log.Fatalf("Error connecting to db: %v", err)
		return
	}
	defer db.Close()

	MigrateDB()

	service.Init()

	profile.RegisterProfileHandler(service.Server(), &handler.ProfileService{
		Auth: auth.NewAuthService("senonerk.sup.srv.auth", service.Client()),
	})

	micro.RegisterSubscriber("sup.auth.NewUser", service.Server(), new(handler.NewUserSubscriber))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

// MigrateDB creates tables in db
func MigrateDB() {
	db.D().AutoMigrate(models.Profile{})
}
