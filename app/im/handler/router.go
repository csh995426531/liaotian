package handler

import (
	"github.com/gin-gonic/gin"
)

func InitRouters(middleware ...gin.HandlerFunc) *gin.Engine {

	router := gin.Default()
	router.Use(middleware...)

	userGroup := router.Group("/user")
	{
		//注册
		userGroup.POST("/register", Register)
		//登录
		userGroup.POST("/login", Login)
		//创建用户信息
		userGroup.GET("/info", GetUserInfo)
		//更新用户信息
		userGroup.POST("/info", UpdateUserInfo)
	}

	applicationGroup := router.Group("/application")
	{
		//创建申请单
		applicationGroup.POST("/create", CreateApplication)
		//申请单列表
		applicationGroup.GET("/list", applicationList)
		//申请单信息
		applicationGroup.GET("/info", ApplicationInfo)
		//通过申请
		applicationGroup.GET("/pass", PassApplication)
		//拒绝申请
		applicationGroup.GET("/reject", RejectApplication)
		//回复申请
		applicationGroup.GET("/reply", ReplyApplication)
	}

	friendGroup := router.Group("/friend")
	{
		//朋友列表
		friendGroup.GET("/list", FriendList)
		//朋友信息
		friendGroup.GET("/info", FriendInfo)
	}

	return router
}

func AddMiddleware(router *gin.Engine, middleware gin.HandlerFunc) *gin.Engine {

	router.Use(middleware)
	return router
}
