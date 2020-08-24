package handler

import "github.com/gin-gonic/gin"

func InitRouters() *gin.Engine {
	router := gin.Default()

	router.POST("login", Login)

	return router
}