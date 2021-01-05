package main

import (
	"github.com/SkyAPM/go2sky"
	"github.com/SkyAPM/go2sky/reporter"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/kubernetes"
	"liaotian/domain/user/repository"
	"liaotian/middlewares/logger/zap"
	"liaotian/middlewares/wrapper/skywalking/micro2sky"
	"time"
)

func main() {
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

	service := micro.NewService(
		micro.Name("domain.friend.service"),
		micro.Registry(kubernetes.NewRegistry()),
		micro.RegisterTTL(time.Second*15),
		micro.Version("latest"),
		micro.WrapHandler(micro2sky.NewHandlerWrapper(tracer, "domain.user.service")),
	)

	service.Init()

	proto.Registe
}
