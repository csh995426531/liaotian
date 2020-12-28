package handler

import (
	"github.com/SkyAPM/go2sky"
	"github.com/SkyAPM/go2sky/reporter"
	"github.com/gin-gonic/gin"
	"liaotian/middlewares/logger/zap"
	"liaotian/middlewares/wrapper/skywalking/gin2micro"
)

func InitRouters() *gin.Engine {

	report, err := reporter.NewGRPCReporter("oap.skywalking.svc.cluster.local:11800")
	if err != nil {
		zap.SugarLogger.Fatalf("创建 grpc reporter 失败，error: %v", err)
	}
	tracer, err := go2sky.NewTracer("app-im", go2sky.WithReporter(report))
	if err != nil {
		zap.SugarLogger.Fatalf("创建 tracer 失败，error: %v", err)
	} else {
		zap.ZapLogger.Info("创建 trace oap.skywalking:11800 - app-im 成功")
	}

	router := gin.Default()
	router.Use(gin2micro.Middleware(router, tracer))

	userGroup := router.Group("/user")
	{
		userGroup.POST("/login", Login)
		userGroup.POST("/register", Register)
		userGroup.GET("/info", GetUserInfo)
		userGroup.POST("/info", UpdateUserInfo)
	}

	return router
}