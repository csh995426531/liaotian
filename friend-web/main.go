package main

import (
	"liaotian/friend-web/handler"
	"log"
	"os"

	"github.com/micro/cli"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/kubernetes"
)

func main() {

	ginRouter := handler.InitRouters()

	// create new web service
	service := web.NewService(
		web.Registry(kubernetes.NewRegistry()),
		web.Handler(ginRouter),
		web.Name("friend.web.friend"),
		web.Version("latest"),
		web.Address(os.Getenv("SERVICE_PORT")),
	)

	// initialise service
	if err := service.Init(
		web.Action(func(context *cli.Context) {
			handler.Init()
		}),
	); err != nil {
		log.Fatal(err)
	}
	// run service
	if err := service.Run(); err != nil {
		log.Fatal()
		log.Fatal(err)
	}

}
