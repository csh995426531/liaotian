package main

import (
	"liaotian/user-service/config"
	"liaotian/user-service/handler"
	"liaotian/user-service/repository"
	"time"

	"github.com/SkyAPM/go2sky"
	sky2micro "github.com/SkyAPM/go2sky-plugins/micro"
	"github.com/SkyAPM/go2sky/reporter"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/registry/kubernetes"

	proto "liaotian/user-service/proto/user"
)

func main() {
	config.Init()

	report, err := reporter.NewGRPCReporter(config.SkywalkingConfig.Url)
	if err != nil {
		log.Fatalf("crate grpc reporter error: %v \n", err)
	}
	tracer, err := go2sky.NewTracer("user-service", go2sky.WithReporter(report))
	if err != nil {
		log.Fatalf("crate tracer error: %v \n", err)
	} else {
		log.Infof("create trace oap.skywalking:11800 - user-service")
	}

	// 新建服务
	service := micro.NewService(
		micro.Name("user.service.user"),
		micro.Registry(kubernetes.NewRegistry()), //注册到Kubernetes
		micro.Version("latest"),
		micro.RegisterTTL(time.Second*15),
		micro.WrapHandler(sky2micro.NewHandlerWrapper(tracer, "user-service")),
	)

	// 服务初始化
	service.Init()

	// 注册服务
	_ = proto.RegisterUserHandler(service.Server(), handler.New(repository.Init()))

	// 启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
