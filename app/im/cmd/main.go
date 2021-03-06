package main

import (
	"github.com/SkyAPM/go2sky"
	"github.com/SkyAPM/go2sky/reporter"
	_ "github.com/micro/go-micro/broker"
	_ "github.com/micro/go-micro/broker/nats"
	"github.com/micro/go-plugins/broker/grpc"
	"liaotian/app/im/event"
	"liaotian/app/im/handler"
	"liaotian/middlewares/logger/zap"
	"liaotian/middlewares/wrapper/skywalking/gin2micro"
	"os"

	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/kubernetes"
)

/**
包入口
*/
func main() {

	zap.InitLogger()

	//初始化路由
	report, err := reporter.NewGRPCReporter("oap.skywalking.svc.cluster.local:11800")
	if err != nil {
		zap.SugarLogger.Fatalf("创建 grpc reporter 失败，error: %v", err)
	}
	tracer, err := go2sky.NewTracer("app-im", go2sky.WithReporter(report))
	if err != nil {
		zap.SugarLogger.Fatalf("创建 tracer 失败，error: %v", err)
	} else {
		zap.ZapLogger.Info("创建 trace oap.skywalking:11800 - app-im success")
	}

	//natsBroker := nats.NewBroker(broker.Addrs("nats-cluster.nats.svc.cluster.local:4222"))
	grpcBroker := grpc.NewBroker()
	grpcBroker.Init()
	event.Init(grpcBroker)

	handler.Init()
	ginRouter := handler.InitRouters()
	ginRouter = handler.AddMiddleware(ginRouter, gin2micro.Middleware(ginRouter, tracer))

	// create new web handler
	service := web.NewService(
		web.Registry(kubernetes.NewRegistry()), //注册到Kubernetes
		web.Name("app.im.service"),
		web.Version("latest"),
		web.Handler(ginRouter),
		web.Address(os.Getenv("SERVICE_PORT")),
	)

	// 服务初始化
	if err := service.Init(); err != nil {
		zap.SugarLogger.Fatalf("服务初始化失败，error: %v", err)
	}

	// run handler
	if err := service.Run(); err != nil {
		zap.SugarLogger.Fatalf("服务启动失败，error: %v", err)
	}
	zap.ZapLogger.Info("服务启动success")
}
