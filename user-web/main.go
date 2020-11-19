package main

import (
	"liaotian/user-web/handler"
	"os"

	"github.com/micro/cli"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/kubernetes"
)

func main() {

	// tracer, err := skywalking.InitTracer("oap.skywalking:11800", "user-service")
	// if err != nil {
	// }

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

	// 新建服务
	// service := micro.NewService(
	// 	micro.Name("user.web.user"),
	// 	micro.Registry(kubernetes.NewRegistry()), //注册到Kubernetes
	// 	micro.Version("latest"),
	// 	micro.Address(os.Getenv("SERVICE_PORT")),
	// 	// micro.RegisterHandler(ginRouter),
	// 	micro.WrapHandler(skywalking.NewHandlerWrapper(tracer, "user-service")),
	// )

	// 服务初始化
	// initialise service
	if err := service.Init(
		web.Action(func(c *cli.Context) {
			handler.Init()
		}),
	); err != nil {
		log.Fatal(err)
	}

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
