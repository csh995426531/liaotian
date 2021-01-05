package handler

import (
	"github.com/gin-gonic/gin"
	"liaotian/app/im/handler/validator"
	ginResult "liaotian/middlewares/common-result/gin"
	"liaotian/middlewares/logger/zap"
	"liaotian/middlewares/tool"
	"net/http"
)

//登录
func Login(ctx *gin.Context) {

	loginValidator := &validator.LoginValidator{}
	req, err := validator.Bind(ctx, loginValidator)

	if err != nil {
		ginResult.Failed(ctx, http.StatusBadRequest, err.Error())
		return
	}

	res, err := domainUser.CheckUserPwd(ctx.Request.Context(), &req)

	if err != nil {
		zap.SugarLogger.Errorf("domainUser.CheckUserPwd error: %+v", err)
		ginResult.Failed(ctx, http.StatusInternalServerError, "上游服务异常")
		return
	}
	if res.Code != http.StatusOK {
		ginResult.Failed(ctx, tool.Int64ToInt(res.Code), res.Message)
		return
	}

	ginResult.Success(ctx, http.StatusOK, res.Data)
}

//注册
func Register(ctx *gin.Context) {

	registerValidator := &validator.RegisterValidator{}
	req, err := validator.Bind(ctx, registerValidator)
	if err != nil {
		ginResult.Failed(ctx, http.StatusBadRequest, err.Error())
		return
	}

	res, err := domainUser.CreateUserInfo(ctx.Request.Context(), &req)
	if err != nil {
		zap.SugarLogger.Errorf("domainUser.CreateUserInfo error: %+v", err)
		ginResult.Failed(ctx, http.StatusInternalServerError, "上游服务异常")
		return
	}

	if res.Code != http.StatusCreated {
		ginResult.Failed(ctx, tool.Int64ToInt(res.Code), res.Message)
		return
	}

	ginResult.Success(ctx, http.StatusCreated, res.Data)
	return
}

//获取用户信息
func GetUserInfo(ctx *gin.Context) {

	getUserInfoValidator := &validator.GetUserInfoValidator{}
	req, err := validator.Bind(ctx, getUserInfoValidator)
	if err != nil {
		ginResult.Failed(ctx, http.StatusBadRequest, err.Error())
		return
	}
	res, err := domainUser.GetUserInfo(ctx.Request.Context(), &req)
	if err != nil {
		zap.SugarLogger.Errorf("domainUser.GetUserInfo error: %+v", err)
		ginResult.Failed(ctx, http.StatusInternalServerError, "上游服务异常")
		return
	}
	if res.Code != http.StatusOK {
		ginResult.Failed(ctx, tool.Int64ToInt(res.Code), res.Message)
		return
	}

	ginResult.Success(ctx, http.StatusOK, res.Data)
	return
}

//更新用户信息
func UpdateUserInfo(ctx *gin.Context) {

	updateUserInfoValidator := &validator.UpdateUserInfoValidator{}
	req, err := validator.Bind(ctx, updateUserInfoValidator)
	if err != nil {
		ginResult.Failed(ctx, http.StatusBadRequest, err.Error())
		return
	}
	res, err := domainUser.UpdateUserInfo(ctx.Request.Context(), &req)
	if err != nil {
		zap.SugarLogger.Errorf("domainUser.UpdateUserInfo error: %+v", err)
		ginResult.Failed(ctx, http.StatusInternalServerError, "上游服务异常")
		return
	}

	if res.Code != http.StatusOK {
		ginResult.Failed(ctx, tool.Int64ToInt(res.Code), res.Message)
		return
	}

	ginResult.Success(ctx, http.StatusOK, res.Data)
	return
}
