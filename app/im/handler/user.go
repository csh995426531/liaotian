package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	userService "liaotian/domain/user/proto"
	ginResult "liaotian/middlewares/common-result/gin"
	"liaotian/middlewares/logger/zap"
	"net/http"
)

func Login(ctx *gin.Context) {

	var request userService.Request
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ginResult.Failed(ctx, http.StatusBadRequest, fmt.Sprintf("传参错误：%+v", err))
		return
	}

	res, err := domainUser.CheckUserPwd(ctx.Request.Context(), &request)

	if err != nil {
		zap.SugarLogger.Errorf("domainUser.CheckUserPwd error: %+v", err)
		ginResult.Failed(ctx, http.StatusInternalServerError, "上游服务异常")
		return
	}
	if res.Code != http.StatusOK {
		ginResult.Failed(ctx, http.StatusInternalServerError, res.Message)
		return
	}

	ginResult.Success(ctx, http.StatusOK, res.Data)
}

func Register(ctx *gin.Context) {
	var result userService.Request

	err := ctx.ShouldBindJSON(&result)
	if err != nil {
		ginResult.Failed(ctx, http.StatusBadRequest, fmt.Sprintf("传参错误：%+v", err))
		return
	}

	res, err := domainUser.CreateUserInfo(ctx.Request.Context(), &result)
	if err != nil {
		zap.SugarLogger.Errorf("domainUser.CreateUserInfo error: %+v", err)
		ginResult.Failed(ctx, http.StatusInternalServerError, "上游服务异常")
		return
	}

	if res.Code != http.StatusOK {
		ginResult.Failed(ctx, http.StatusInternalServerError, res.Message)
		return
	}

	ginResult.Success(ctx, http.StatusOK, res.Data)
}

func GetUserInfo(ctx *gin.Context) {

	var result userService.Request

	err := ctx.ShouldBindJSON(&result)
	if err != nil {
		ginResult.Failed(ctx, http.StatusBadRequest, fmt.Sprintf("传参错误：%+v", err))
		return
	}

	res, err := domainUser.GetUserInfo(ctx.Request.Context(), &result)
	if err != nil {
		zap.SugarLogger.Errorf("domainUser.GetUserInfo error: %+v", err)
		ginResult.Failed(ctx, http.StatusInternalServerError, "上游服务异常")
		return
	}
	if res.Code != http.StatusOK {
		ginResult.Failed(ctx, http.StatusInternalServerError, res.Message)
		return
	}

	ginResult.Success(ctx, http.StatusOK, res.Data)
}

func UpdateUserInfo(ctx *gin.Context) {

	var result userService.Request

	err := ctx.ShouldBindJSON(&result)
	if err != nil {
		ginResult.Failed(ctx, http.StatusBadRequest, fmt.Sprintf("传参错误：%+v", err))
		return
	}

	res, err := domainUser.UpdateUserInfo(ctx.Request.Context(), &result)
	if err != nil {
		ginResult.Failed(ctx, http.StatusInternalServerError, "上游服务异常")
		return
	}

	if res.Code != http.StatusOK {
		zap.SugarLogger.Errorf("domainUser.UpdateUserInfo error: %+v", err)
		ginResult.Failed(ctx, http.StatusInternalServerError, res.Message)
		return
	}

	ginResult.Success(ctx, http.StatusOK, res.Data)
}