package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/registry/kubernetes/v2"
	"liaotian/friend-service/config"
	"liaotian/friend-service/handler"
	"liaotian/friend-service/repository"

	"time"

	proto "liaotian/friend-service/proto/friend"
)

func main() {
	
	config.Init()

	// New Service
	service := micro.NewService(
		micro.Name("friend.service.friend"),
		micro.Registry(kubernetes.NewRegistry()),
		micro.Version("latest"),
		micro.RegisterTTL(time.Second*15),
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
