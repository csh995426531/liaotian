package main

import (
	"liaotian/app/im/handler"
	"liaotian/middlewares/logger/zap"
	"os"

	"github.com/micro/cli"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/kubernetes"
)

func main() {
	// 日志
	zap.InitLogger()

	//初始化路由
	ginRouter := handler.InitRouters()

	// create new web handler
	service := web.NewService(
		web.Registry(kubernetes.NewRegistry()), //注册到Kubernetes
		web.Name("user.web.user"),
		web.Version("latest"),
		web.Handler(ginRouter),
		web.Address(os.Getenv("SERVICE_PORT")),
	)

	// 服务初始化
	if err := service.Init(
		web.Action(func(c *cli.Context) {
			handler.Init()
		}),
	); err != nil {
		zap.SugarLogger.Fatalf("服务初始化失败，error: %v", err)
	}

	// run handler
	if err := service.Run(); err != nil {
		zap.SugarLogger.Fatalf("服务启动失败，error: %v", err)
	}
	zap.ZapLogger.Info("服务启动成功")
}
