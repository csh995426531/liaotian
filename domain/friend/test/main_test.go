package test

import (
	"fmt"
	"github.com/micro/go-micro"
	"liaotian/domain/friend/handler"
	"liaotian/domain/friend/proto"
	"liaotian/domain/friend/repository"
	"liaotian/middlewares/logger/zap"
	"testing"
	"time"
)

func TestMain(m *testing.M) {

	zap.InitLogger()

	db, mockDb := repository.NewMockDb()
	repository.Init(db, mockDb)

	// 新建服务
	service := micro.NewService(
		micro.Name("domain.friend.service"),
		micro.Version("latest"),
		micro.RegisterTTL(time.Second*15),
	)

	// 注册服务
	_ = proto.RegisterFriendHandler(service.Server(), handler.Init())

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
	t.Run("CreateApplicationInfo", CreateApplicationInfo)
	t.Run("GetApplicationInfo", GetApplicationInfo)
	t.Run("PassApplicationInfo", PassApplicationInfo)
	t.Run("RejectApplicationInfo", RejectApplicationInfo)
	t.Run("GetApplicationList", GetApplicationList)
	t.Run("CreateApplicationSay", CreateApplicationSay)
	t.Run("GetFriendList", GetFriendList)
	t.Run("DeleteFriendInfo", DeleteFriendInfo)
}
