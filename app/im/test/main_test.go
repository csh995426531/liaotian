package test

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/micro/go-micro/web"
	"liaotian/app/im/handler"
	authService "liaotian/domain/auth/proto"
	"liaotian/middlewares/logger/zap"
	"testing"
	"time"
)

/**
测试入口
*/
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

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := NewMockAuthService(ctrl)
	handler.AuthDomain(service)
	service.EXPECT().Parse(gomock.Any(), gomock.Any()).Return(&authService.ParseResponse{
		Message: "success",
		Data: &authService.User{
			UserId: 1,
			Name:   "张三",
		},
	}, nil)

	t.Run("register", register)
	t.Run("login", login)
	t.Run("getUserInfo", getUserInfo)
	t.Run("updateUserInfo", updateUserInfo)
	t.Run("createApplication", createApplication)
	t.Run("applicationList", applicationList)
	t.Run("applicationInfo", applicationInfo)
	t.Run("passApplication", passApplication)
	t.Run("rejectApplication", rejectApplication)
	t.Run("replyApplication", replyApplication)
}
