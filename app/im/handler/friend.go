package handler

import (
	"github.com/gin-gonic/gin"
	"liaotian/app/im/handler/validator"
	friendService "liaotian/domain/friend/proto"
	ginResult "liaotian/middlewares/common-result/gin"
	"liaotian/middlewares/logger/zap"
	"net/http"
)

/**
好友控制器
*/
//好友列表
func FriendList(ctx *gin.Context) {
	friendListValidator := &validator.FriendListValidator{}
	req := &friendService.GetFriendListRequest{}

	err := validator.Bind(ctx, friendListValidator, req)
	if err != nil {
		ginResult.Failed(ctx, http.StatusBadRequest, err.Error())
		return
	}

	res, err := domainFriend.GetFriendList(ctx.Request.Context(), req)
	if err != nil {
		zap.SugarLogger.Errorf("domainFriend.GetFriendList error: %v", err)
		ginResult.Failed(ctx, http.StatusInternalServerError, "上游服务异常")
		return
	}

	// 获取好友用户信息

	ginResult.Success(ctx, http.StatusOK, res.Data)
}

//删除好友
func DeleteFriendInfo(ctx *gin.Context) {

	deleteFriendInfoValidator := &validator.DeleteFriendInfoValidator{}
	req := &friendService.DeleteFriendInfoRequest{}

	err := validator.Bind(ctx, deleteFriendInfoValidator, req)
	if err != nil {
		ginResult.Failed(ctx, http.StatusBadRequest, err.Error())
		return
	}
	res, err := domainFriend.DeleteFriendInfo(ctx.Request.Context(), req)
	if err != nil {
		zap.SugarLogger.Errorf("domainFriend.DeleteFriendInfo error: %v", err)
		ginResult.Failed(ctx, http.StatusInternalServerError, "上游服务异常")
		return
	}
	ginResult.Success(ctx, http.StatusOK, res.Ok)
}

//好友信息
func FriendInfo(ctx *gin.Context) {

	//friendInfoValidator := &validator.FriendInfoValidator{}
	//friendService.g

	return
}
