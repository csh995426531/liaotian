package handler

import (
	"github.com/gin-gonic/gin"
)

func InitRouters(middleware ...gin.HandlerFunc) *gin.Engine {

	router := gin.Default()
	router.Use(middleware...)

	userGroup := router.Group("/user")
	{
		userGroup.POST("/register", Register)
		userGroup.POST("/login", Login)
		userGroup.GET("/info", GetUserInfo)
		userGroup.POST("/info", UpdateUserInfo)
	}

	return router
}

func AddMiddleware(router *gin.Engine, middleware gin.HandlerFunc) *gin.Engine {

	router.Use(middleware)
	return router
}
