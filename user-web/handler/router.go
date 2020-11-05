package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/util/log"
	"github.com/opentracing/opentracing-go"
	"liaotian/plugins/skywalking"
	"liaotian/plugins/wrapper/tracer/opentracing/gin2micro"
)

func InitRouters() *gin.Engine {
	gin2micro.SetSamplingFrequency(50)
	t, io, err := skywalking.NewTracer("user.web.user", "")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)


	router := gin.Default()
	router.Use(gin2micro.TracerWrapper)
	router.POST("login", Login)

	return router
}