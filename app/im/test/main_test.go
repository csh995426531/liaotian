package test

import (
	"fmt"
	"github.com/micro/go-micro/web"
	"liaotian/app/im/handler"
	"liaotian/middlewares/logger/zap"
	"testing"
	"time"
)

func TestMain(m *testing.M) {

	zap.InitLogger()
	//translate.Init()

	//初始化路由
	ginRouter := handler.InitRouters()

	// create new web handler
	service := web.NewService(
		web.Name("app.im.service"),
		web.Version("latest"),
		web.Handler(ginRouter),
		web.Address(":18282"),
	)
	handler.UserDomain(new(testService))

	// run handler
	go func() {
		if err := service.Run(); err != nil {
			panic(fmt.Sprintf("服务启动失败，error: %v", err))
		}
	}()

	fmt.Println("服务启动成功")
	time.Sleep(time.Second * 1)
	m.Run()
}

func TestStart(t *testing.T) {
	t.Run("Register", Register)
	t.Run("Login", Login)
	t.Run("GetUserInfo", GetUserInfo)
	t.Run("UpdateUserInfo", UpdateUserInfo)
}
