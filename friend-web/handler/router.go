package handler

import "github.com/gin-gonic/gin"

func InitRouters() *gin.Engine {
	router := gin.Default()

	router.POST("add", Add)
	router.DELETE("del", Del)
	router.GET("list", List)
	router.GET("info", Info)

	return router
}