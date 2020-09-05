package main

import (
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-plugins/registry/kubernetes/v2"
	"liaotian/user-web/handler"
	"os"
)

func main() {

	//初始化路由
	ginRouter := handler.InitRouters()

	// create new web service
	service := web.NewService(
		web.Registry(kubernetes.NewRegistry()), //注册到Kubernetes
		web.Name("user.web.user"),
		web.Version("latest"),
		web.Handler(ginRouter),
		web.Address(os.Getenv("SERVICE_PORT")),
	)


	// initialise service
	if err := service.Init(
		web.Action(func(c *cli.Context) {
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
