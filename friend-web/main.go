package main

import (
        "github.com/micro/cli/v2"
        "github.com/micro/go-micro/v2/logger"
        "github.com/micro/go-plugins/registry/kubernetes/v2"
        "github.com/micro/go-micro/v2/web"
        "liaotian/friend-web/handler"
        "os"
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
                logger.Fatal(err)
        }
	// run service
        if err := service.Run(); err != nil {
                logger.Fatal(err)
        }
}
