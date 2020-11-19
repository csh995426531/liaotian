package handler

import (
	"github.com/gin-gonic/gin"
	// "liaotian/plugins/wrapper/tracer/opentracing/gin2micro"
)

func InitRouters() *gin.Engine {
	// gin2micro.SetSamplingFrequency(50)

	router := gin.Default()
	// router.Use(gin2micro.TracerWrapper)
	router.POST("login", Login)

	return router
}
