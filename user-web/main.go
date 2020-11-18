package main

import (
	"liaotian/user-web/handler"
	"os"

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-plugins/registry/kubernetes/v2"
)

func main() {

	report, err = reporter.NewGRPCReporter("oap.skywalking:11800")
	if err != nil {
		log.Fatalf("crate grpc reporter error: %v \n", err)
	}
	tracer, err := go2sky.NewTracer(, go2sky.WithReporter(report))
	if err != nil {
		log.Fatalf("crate tracer error: %v \n", err)
	}
	
	//初始化路由
	ginRouter := handler.InitRouters()

	create new web service
	service := web.NewService(
		web.Registry(kubernetes.NewRegistry()), //注册到Kubernetes
		web.Name("user.web.user"),
		web.Version("latest"),
		web.Handler(ginRouter),
		web.Address(os.Getenv("SERVICE_PORT")),
	)

	// 新建服务
	service := micro.NewService(
		micro.Name("user.web.user"),
		micro.Registry(kubernetes.NewRegistry()), //注册到Kubernetes
		micro.Version("latest"),
		micro.Address(os.Getenv("SERVICE_PORT")),
		micro.WrapHandler(skywalking.NewHandlerWrapper(tracer, "user-service")),
	)

	// 服务初始化
	// initialise service
	if err := service.Init(
		micro.Action(func(c *cli.Context) {
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
