package main

import (
	"liaotian/domain/user/handler"
	"liaotian/domain/user/proto"
	"liaotian/domain/user/repository"
	"liaotian/middlewares/logger/zap"
	"liaotian/middlewares/wrapper/skywalking/micro2sky"
	"time"

	"github.com/SkyAPM/go2sky"
	"github.com/SkyAPM/go2sky/reporter"

	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/kubernetes"
)

/**
用户领域服务入口
*/

func main() {
	//config.Init()
	zap.InitLogger()
	repository.Init(repository.NewDb(), nil)

	report, err := reporter.NewGRPCReporter("oap.skywalking.svc.cluster.local:11800")
	if err != nil {
		zap.SugarLogger.Fatalf("创建grpc reporter失败，error: %v", err)
	}
	tracer, err := go2sky.NewTracer("domain.user.service", go2sky.WithReporter(report))
	if err != nil {
		zap.SugarLogger.Fatalf("创建tracer失败，error: %v", err)
	} else {
		zap.ZapLogger.Info("创建 trace oap.skywalking:11800 - domain.user.service 成功")
	}

	// 新建服务
	service := micro.NewService(
		micro.Name("domain.user.service"),
		micro.Registry(kubernetes.NewRegistry()), //注册到Kubernetes
		micro.Version("latest"),
		micro.RegisterTTL(time.Second*15),
		micro.WrapHandler(micro2sky.NewHandlerWrapper(tracer, "domain.user.service")),
	)

	// 服务初始化
	service.Init()

	// 注册服务
	_ = proto.RegisterUserHandler(service.Server(), handler.Init())

	// 启动服务
	if err := service.Run(); err != nil {
		zap.SugarLogger.Fatalf("服务启动失败，error: %v", err)
	}
	zap.ZapLogger.Info("服务启动成功")
}
