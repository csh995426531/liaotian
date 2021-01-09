package handler

import (
	"github.com/gin-gonic/gin"
	"liaotian/app/im/handler/validator"
	friendService "liaotian/domain/friend/proto"
	ginResult "liaotian/middlewares/common-result/gin"
	"liaotian/middlewares/logger/zap"
	"net/http"
)

//创建申请单
func CreateApplication(ctx *gin.Context) {
	createApplicationValidator := &validator.CreateApplicationValidator{}
	req := friendService.CreateApplicationRequest{}
	err := validator.Bind(ctx, createApplicationValidator, &req)
	if err != nil {
		ginResult.Failed(ctx, http.StatusBadRequest, err.Error())
		return
	}

	res, err := domainFriend.CreateApplicationInfo(ctx.Request.Context(), &req)
	if err != nil {
		zap.SugarLogger.Errorf("domainFriend.CreateApplicationInfo error: %+v", err)
		ginResult.Failed(ctx, http.StatusInternalServerError, "上游服务异常")
		return
	}

	ginResult.Success(ctx, http.StatusOK, res.Data)
}

//申请单列表
func applicationList(ctx *gin.Context) {
	applicationListValidator := &validator.ApplicationListValidator{}
	req := &friendService.GetApplicationListRequest{}

	err := validator.Bind(ctx, applicationListValidator, req)
	if err != nil {
		ginResult.Failed(ctx, http.StatusBadRequest, err.Error())
		return
	}

	res, err := domainFriend.GetApplicationList(ctx.Request.Context(), req)
	if err != nil {
		zap.SugarLogger.Errorf("domainFriend.GetApplicationList error: %+v", err)
		ginResult.Failed(ctx, http.StatusInternalServerError, "上游服务异常")
		return
	}

	//获取用户信息

	ginResult.Success(ctx, http.StatusOK, res.Data)
}

//申请单信息
func ApplicationInfo(ctx *gin.Context) {
	applicationInfoValidator := &validator.ApplicationInfoValidator{}
	req := &friendService.GetApplicationRequest{}

	err := validator.Bind(ctx, applicationInfoValidator, req)
	if err != nil {
		ginResult.Failed(ctx, http.StatusBadRequest, err.Error())
		return
	}

	res, err := domainFriend.GetApplicationInfo(ctx.Request.Context(), req)
	if err != nil {
		zap.SugarLogger.Errorf("domainFriend.GetApplicationInfo error: %v", err)
		ginResult.Failed(ctx, http.StatusInternalServerError, "上游服务异常")
		return
	}

	ginResult.Success(ctx, http.StatusOK, res.Data)
}

//通过申请
func PassApplication(ctx *gin.Context) {
	passApplicationValidator := &validator.PassApplicationValidator{}
	req := &friendService.PassApplicationInfoRequest{}

	err := validator.Bind(ctx, passApplicationValidator, req)
	if err != nil {
		ginResult.Failed(ctx, http.StatusBadRequest, err.Error())
		return
	}

	res, err := domainFriend.PassApplicationInfo(ctx.Request.Context(), req)
	if err != nil {
		zap.SugarLogger.Errorf("domainFriend.PassApplicationInfo error: %v", err)
		ginResult.Failed(ctx, http.StatusInternalServerError, "上游服务异常")
		return
	}

	ginResult.Success(ctx, http.StatusOK, res.Ok)
}

//拒绝申请
func RejectApplication(ctx *gin.Context) {
	rejectApplicationValidator := &validator.RejectApplicationValidator{}
	req := &friendService.RejectApplicationInfoRequest{}

	err := validator.Bind(ctx, rejectApplicationValidator, req)
	if err != nil {
		ginResult.Failed(ctx, http.StatusBadRequest, err.Error())
		return
	}

	res, err := domainFriend.RejectApplicationInfo(ctx.Request.Context(), req)
	if err != nil {
		zap.SugarLogger.Errorf("domainFriend.RejectApplicationInfo error: %v", err)
		ginResult.Failed(ctx, http.StatusInternalServerError, "上游服务异常")
		return
	}

	ginResult.Success(ctx, http.StatusOK, res.Ok)
}

//回复申请
func ReplyApplication(ctx *gin.Context) {
	replyApplicationValidator := &validator.ReplyApplicationValidator{}
	req := &friendService.CreateApplicationSayRequest{}

	err := validator.Bind(ctx, replyApplicationValidator, req)
	if err != nil {
		ginResult.Failed(ctx, http.StatusBadRequest, err.Error())
		return
	}

	resp, err := domainFriend.CreateApplicationSay(ctx.Request.Context(), req)
	if err != nil {
		zap.SugarLogger.Errorf("domainFriend.CreateApplicationSay error: %v", err)
		ginResult.Failed(ctx, http.StatusInternalServerError, "上游服务异常")
		return
	}

	ginResult.Success(ctx, http.StatusOK, resp.Data)
}
