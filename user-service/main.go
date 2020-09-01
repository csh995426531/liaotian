package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-plugins/registry/kubernetes/v2"
	"liaotian/user-service/config"
	"liaotian/user-service/handler"
	"liaotian/user-service/repository"
	"time"

	proto "liaotian/user-service/proto/user"
)



func main() {
	config.Init()
	// 新建服务
	service := micro.NewService(
		micro.Name("user.service.user"),
		micro.Registry(kubernetes.NewRegistry()),//注册到Kubernetes
		micro.Version("latest"),
		micro.RegisterTTL(time.Second*15),
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
