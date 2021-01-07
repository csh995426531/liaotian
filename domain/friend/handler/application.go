package handler

import (
	"context"
	"errors"
	"liaotian/domain/friend/entity"
	"liaotian/domain/friend/proto"
	"liaotian/middlewares/logger/zap"
)

func (h *Handler) CreateApplicationInfo(ctx context.Context, request *proto.CreateApplicationRequest, response *proto.ApplicationResponse) error {
	if request.SenderId == 0 || request.ReceiverId == 0 {
		return ErrorBadRequest
	}
	application, err := h.ApplicationEntity.CreateApplicationInfo(request.SenderId, request.ReceiverId)
	if err != nil {
		zap.ZapLogger.Error(err.Error())
		return ErrorInternalServerError(err)
	}
	response.Message = "success"
	response.Data = &proto.Application{
		Id:         application.Id,
		SenderId:   application.SenderId,
		ReceiverId: application.ReceiverId,
	}

	return nil
}

func (h *Handler) GetApplicationInfo(ctx context.Context, request *proto.GetApplicationRequest, response *proto.ApplicationResponse) error {
	if request.Id == 0 {
		return ErrorBadRequest
	}

	application, err := h.ApplicationEntity.GetApplicationInfo(request.Id)
	if err != nil {
		zap.ZapLogger.Error(err.Error())
		return ErrorInternalServerError(err)
	}
	if application.Id == 0 {
		return ErrorApplicationNotFound
	}

	response.Message = "success"
	response.Data = &proto.Application{
		Id:         application.Id,
		SenderId:   application.SenderId,
		ReceiverId: application.ReceiverId,
	}
	for _, say := range application.SayList {
		response.Data.SayList = append(response.Data.SayList, &proto.ApplicationSay{
			Id:       say.Id,
			SenderId: say.SenderId,
			Content:  say.Content,
		})
	}

	return nil
}

func (h *Handler) PassApplicationInfo(ctx context.Context, request *proto.PassApplicationInfoRequest, response *proto.Response) error {
	if request.Id == 0 {
		return ErrorBadRequest
	}

	application, err := h.ApplicationEntity.GetApplicationInfo(request.Id)
	if err != nil {
		zap.ZapLogger.Error(err.Error())
		return ErrorInternalServerError(err)
	}
	if application.Id == 0 {
		return ErrorApplicationNotFound
	}
	if application.Status != entity.StatusWait {
		return ErrorApplicationStatusWrong
	}

	ok, err := h.ApplicationEntity.UpdateApplicationInfoStatus(request.Id, entity.StatusPass)
	if err != nil {
		zap.ZapLogger.Error(err.Error())
		return ErrorInternalServerError(err)
	}

	friend, err := h.FriendEntity.CreateFriendInfo(application.SenderId, application.ReceiverId)
	if err != nil {
		zap.ZapLogger.Error(err.Error())
		return ErrorInternalServerError(err)
	}
	if friend.Id == 0 {
		return ErrorInternalServerError(errors.New("创建用户失败"))
	}

	response.Message = "success"
	response.Ok = ok
	return nil
}

func (h *Handler) RejectApplicationInfo(ctx context.Context, request *proto.RejectApplicationInfoRequest, response *proto.Response) error {
	if request.Id == 0 {
		return ErrorBadRequest
	}

	application, err := h.ApplicationEntity.GetApplicationInfo(request.Id)
	if err != nil {
		zap.ZapLogger.Error(err.Error())
		return ErrorInternalServerError(err)
	}
	if application.Id == 0 {
		return ErrorApplicationNotFound
	}
	if application.Status != entity.StatusWait {
		return ErrorApplicationStatusWrong
	}

	ok, err := h.ApplicationEntity.UpdateApplicationInfoStatus(request.Id, entity.StatusReject)
	if err != nil {
		zap.ZapLogger.Error(err.Error())
		return ErrorInternalServerError(err)
	}

	response.Message = "success"
	response.Ok = ok
	return nil
}

func (h *Handler) GetApplicationList(ctx context.Context, request *proto.GetApplicationListRequest, response *proto.GetApplicationListResponse) error {
	if request.UserId == 0 {
		return ErrorBadRequest
	}
	list, err := h.ApplicationEntity.GetApplicationList(request.UserId)
	if err != nil {
		zap.ZapLogger.Error(err.Error())
		return ErrorInternalServerError(err)
	}

	response.Message = "success"
	for _, application := range list {
		temp := &proto.Application{
			Id:         application.Id,
			SenderId:   application.SenderId,
			ReceiverId: application.ReceiverId,
		}
		for _, say := range application.SayList {
			temp.SayList = append(temp.SayList, &proto.ApplicationSay{
				Id:       say.Id,
				SenderId: say.SenderId,
				Content:  say.Content,
			})
		}
		response.Data = append(response.Data, temp)
	}

	return nil
}

func (h *Handler) CreateApplicationSay(ctx context.Context, request *proto.CreateApplicationSayRequest, response *proto.CreateApplicationSayResponse) error {
	if request.ApplicationId == 0 || request.SenderId == 0 || request.Content == "" {
		return ErrorBadRequest
	}

	application, err := h.ApplicationEntity.GetApplicationInfo(request.ApplicationId)
	if err != nil {
		zap.ZapLogger.Error(err.Error())
		return ErrorInternalServerError(err)
	}
	if application.Id == 0 {
		return ErrorApplicationNotFound
	}
	if application.Status != entity.StatusWait {
		return ErrorApplicationStatusWrong
	}

	say, err := h.ApplicationSayEntity.CreateApplicationSay(request.ApplicationId, request.SenderId, request.Content)
	if err != nil {
		zap.ZapLogger.Error(err.Error())
		return ErrorInternalServerError(err)
	}

	response.Message = "success"
	response.Data = &proto.ApplicationSay{
		Id:       say.Id,
		SenderId: say.SenderId,
		Content:  say.Content,
	}

	return nil
}
