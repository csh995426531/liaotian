package handler

import (
	"github.com/gin-gonic/gin"
	ginWrapper "liaotian/middlewares/wrapper/gin"
)

/**
路由
*/
// 初始化路由
func InitRouters(middleware ...gin.HandlerFunc) *gin.Engine {

	router := gin.Default()
	router.Use(middleware...)

	userGroup := router.Group("/user")
	{
		//注册
		userGroup.POST("/register", Register)
		//登录
		userGroup.POST("/login", Login)
		//获取用户信息
		userGroup.GET("/info", ginWrapper.AuthMiddleware(&domainAuth), GetUserInfo)
		//更新用户信息
		userGroup.POST("/info", ginWrapper.AuthMiddleware(&domainAuth), UpdateUserInfo)
	}

	applicationGroup := router.Group("/application", ginWrapper.AuthMiddleware(&domainAuth))
	{
		//创建申请单
		applicationGroup.POST("/create", CreateApplication)
		//申请单列表
		applicationGroup.GET("/list", applicationList)
		//申请单信息
		applicationGroup.GET("/info", ApplicationInfo)
		//通过申请
		applicationGroup.POST("/pass", PassApplication)
		//拒绝申请
		applicationGroup.POST("/reject", RejectApplication)
		//回复申请
		applicationGroup.POST("/reply", ReplyApplication)
	}

	friendGroup := router.Group("/friend", ginWrapper.AuthMiddleware(&domainAuth))
	{
		//好友列表
		friendGroup.GET("/list", FriendList)
		//删除好友
		friendGroup.DELETE("/delete", DeleteFriendInfo)
		//好友信息
		friendGroup.GET("/info", FriendInfo)
	}

	messageGroup := router.Group("/message", ginWrapper.AuthMiddleware(&domainAuth))
	{
		//连接消息
		messageGroup.GET("/connect", Connect)
	}

	return router
}

// 添加中间件
func AddMiddleware(router *gin.Engine, middleware ...gin.HandlerFunc) *gin.Engine {

	router.Use(middleware...)
	return router
}
