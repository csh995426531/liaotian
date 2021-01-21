package main

import (
	_ "github.com/micro/go-micro/broker"
	_ "github.com/micro/go-micro/broker/nats"
	"github.com/micro/go-plugins/broker/grpc"
	"liaotian/domain/message/event"
	"liaotian/domain/message/handler"
	"liaotian/domain/message/proto"
	"liaotian/middlewares/logger/zap"
	"liaotian/middlewares/wrapper/skywalking/micro2sky"
	"time"

	"github.com/SkyAPM/go2sky"
	"github.com/SkyAPM/go2sky/reporter"

	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/kubernetes"
)

/**
消息领域服务入口
*/

func main() {
	//config.Init()
	zap.InitLogger()

	report, err := reporter.NewGRPCReporter("oap.skywalking.svc.cluster.local:11800")
	if err != nil {
		zap.SugarLogger.Fatalf("创建grpc reporter失败，error: %v", err)
	}
	tracer, err := go2sky.NewTracer("domain.message.service", go2sky.WithReporter(report))
	if err != nil {
		zap.SugarLogger.Fatalf("创建tracer失败，error: %v", err)
	} else {
		zap.ZapLogger.Info("创建 trace oap.skywalking:11800 - domain.message.service 成功")
	}

	//natsBroker := nats.NewBroker(broker.Addrs("nats-cluster.nats.svc.cluster.local:4222"))
	grpcBroker := grpc.NewBroker()
	grpcBroker.Init()
	event.Init(grpcBroker)

	// 新建服务
	service := micro.NewService(
		micro.Name("domain.message.service"),
		micro.Registry(kubernetes.NewRegistry()), //注册到Kubernetes
		micro.Version("latest"),
		micro.RegisterTTL(time.Second*15),
		micro.WrapHandler(micro2sky.NewHandlerWrapper(tracer, "domain.message.service")),
		micro.Broker(grpcBroker),
	)

	// 服务初始化
	service.Init()

	// 注册服务
	_ = proto.RegisterMessageHandler(service.Server(), handler.Init())

	// 启动服务
	if err := service.Run(); err != nil {
		zap.SugarLogger.Fatalf("服务启动失败，error: %v", err)
	}
	zap.ZapLogger.Info("服务启动成功")
}
