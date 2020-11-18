package main

import (
	"liaotian/plugins/wrapper/skywalking"
	"liaotian/user-service/config"
	"liaotian/user-service/handler"
	"liaotian/user-service/repository"
	"time"

	"github.com/SkyAPM/go2sky"
	"github.com/SkyAPM/go2sky/reporter"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-plugins/registry/kubernetes/v2"

	proto "liaotian/user-service/proto/user"
)

func main() {
	report, err := reporter.NewGRPCReporter("oap.skywalking:11800")
	if err != nil {
		logger.Fatalf("crate grpc reporter error: %v \n", err)
	}
	tracer, err := go2sky.NewTracer("user-service", go2sky.WithReporter(report))
	if err != nil {
		logger.Fatalf("crate tracer error: %v \n", err)
	}

	config.Init()
	// 新建服务
	service := micro.NewService(
		micro.Name("user.service.user"),
		micro.Registry(kubernetes.NewRegistry()), //注册到Kubernetes
		micro.Version("latest"),
		micro.RegisterTTL(time.Second*15),
		micro.WrapHandler(skywalking.NewHandlerWrapper(tracer, "user-service")),
	)

	// 服务初始化
	service.Init()

	// 注册服务
	_ = proto.RegisterUserHandler(service.Server(), handler.New(repository.Init()))

	// 启动服务
	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}

}
