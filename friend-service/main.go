package main

import (
	"liaotian/friend-service/config"
	"liaotian/friend-service/handler"
	"liaotian/friend-service/repository"

	"github.com/SkyAPM/go2sky"
	sky2micro "github.com/SkyAPM/go2sky-plugins/micro"
	"github.com/SkyAPM/go2sky/reporter"
	"github.com/micro/go-micro/util/log"

	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/kubernetes"

	"time"

	proto "liaotian/friend-service/proto/friend"
)

func main() {

	config.Init()

	report, err := reporter.NewGRPCReporter(config.SkywalkingConfig.Url)
	if err != nil {
		log.Fatalf("crate grpc reporter error: %v \n", err)
	}
	tracer, err := go2sky.NewTracer("friend-service", go2sky.WithReporter(report))
	if err != nil {
		log.Fatalf("crate tracer error: %v \n", err)
	} else {
		log.Infof("create trace oap.skywalking:11800 - friend-service")
	}

	// New Service
	service := micro.NewService(
		micro.Name("friend.service.friend"),
		micro.Registry(kubernetes.NewRegistry()),
		micro.Version("latest"),
		micro.RegisterTTL(time.Second*15),
		micro.WrapHandler(sky2micro.NewHandlerWrapper(tracer, "friend-service")),
	)

	// Initialise service
	service.Init()

	// Register Handler
	_ = proto.RegisterFriendHandler(service.Server(), handler.New(repository.Init()))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
