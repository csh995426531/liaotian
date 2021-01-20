package handler

import (
	"github.com/gin-gonic/gin"
	"liaotian/app/im/handler/validator"
	friendService "liaotian/domain/friend/proto"
	userService "liaotian/domain/user/proto"
	ginResult "liaotian/middlewares/common-result/gin"
	"liaotian/middlewares/logger/zap"
	"liaotian/middlewares/tool"
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
	if len(res.Data) > 0 {
		var ids []int64
		for _, friend := range res.Data {
			ids = append(ids, friend.UserId)
		}
		batchGetUserRequest := &userService.BatchGetUserInfoRequest{
			Ids: ids,
		}

		userRes, err := domainUser.BatchGetUserInfo(ctx.Request.Context(), batchGetUserRequest)

		if err != nil {
			zap.SugarLogger.Errorf("domainUser.BatchGetUserInfo error: %v", err)
			ginResult.Failed(ctx, http.StatusInternalServerError, "上游服务异常")
			return
		}
		type out struct {
			Id      int64  `json:"id"`
			UserId  int64  `json:"user_id"`
			Name    string `json:"name"`
			Avatar  string `json:"avatar"`
			Account string `json:"account"`
		}

		resList := make([]*out, 0)
		for _, user := range userRes.Data {
			for _, friend := range res.Data {
				if friend.UserId == user.Id {
					resList = append(resList, &out{
						Id:      friend.Id,
						UserId:  user.Id,
						Name:    user.Name,
						Avatar:  user.Avatar,
						Account: user.Account,
					})
					break
				}
			}
		}
		ginResult.Success(ctx, http.StatusOK, resList)
	}
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

	friendInfoValidator := &validator.GetUserInfoValidator{}
	req := &userService.Request{}
	err := validator.Bind(ctx, friendInfoValidator, req)
	if err != nil {
		ginResult.Failed(ctx, http.StatusBadRequest, err.Error())
		return
	}

	res, err := domainUser.GetUserInfo(ctx.Request.Context(), req)
	if err != nil {
		zap.SugarLogger.Errorf("domainUser.GetUserInfo error: %v", err)
		ginResult.Failed(ctx, http.StatusInternalServerError, "上游服务异常")
		return
	}
	if res.Code != http.StatusOK {
		ginResult.Failed(ctx, tool.Int64ToInt(res.Code), res.Message)
		return
	}
	ginResult.Success(ctx, http.StatusOK, res.Data)
}
