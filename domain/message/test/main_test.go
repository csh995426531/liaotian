package test

import (
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/broker/grpc"
	"liaotian/domain/message/event"
	"liaotian/domain/message/handler"
	"liaotian/domain/message/proto"
	"liaotian/middlewares/logger/zap"
	"testing"
	"time"
)

func TestMain(m *testing.M) {

	zap.InitLogger()

	grpcBroker := grpc.NewBroker()
	grpcBroker.Init()
	event.Init(grpcBroker)

	// 新建服务
	service := micro.NewService(
		micro.Name("domain.message.service"),
		micro.Version("latest"),
		micro.RegisterTTL(time.Second*15),
		micro.Broker(grpcBroker),
	)

	// 注册服务
	_ = proto.RegisterMessageHandler(service.Server(), handler.Init())

	go func() {
		// 启动服务
		if err := service.Run(); err != nil {
			zap.SugarLogger.Fatalf("服务启动失败，error: %v", err)
		}
	}()

	fmt.Print("服务启动成功")
	time.Sleep(time.Second * 1)
	m.Run()
}

func TestStart(t *testing.T) {
	t.Run("sub", sub)
	t.Run("unSub", unSub)
	t.Run("send", send)
}