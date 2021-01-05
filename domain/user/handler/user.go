package handler

import (
	"context"
	"liaotian/domain/user/proto"
	"net/http"
)

/**
用户领域服务
*/

func (h *Handler) CreateUserInfo(ctx context.Context, request *proto.Request, response *proto.Response) error {

	if request.Account == "" || request.Name == "" || request.Password == "" {
		return ErrorBadRequest
	}

	user, err := h.UserEntity.GetUserInfo(0, "", request.Account)
	if err != nil {
		return ErrorInternalServerError(err)
	}

	if user.Id > 0 {
		return ErrorUserExists
	}

	user, err = h.UserEntity.CreateUserInfo(request.Name, request.Account, request.Password, request.Avatar)
	if err != nil {
		return ErrorInternalServerError(err)
	}

	response.Code = http.StatusCreated
	response.Message = "success"
	response.Data = &proto.User{
		Id:       user.Id,
		Name:     user.Name,
		Password: user.Password,
		Avatar:   user.Avatar,
	}

	return nil
}

func (h *Handler) GetUserInfo(ctx context.Context, request *proto.Request, response *proto.Response) error {

	if request.Account == "" && request.Name == "" && request.Id == 0 {
		return ErrorBadRequest
	}

	user, err := h.UserEntity.GetUserInfo(request.Id, request.Name, request.Account)
	if err != nil {
		return ErrorInternalServerError(err)
	}

	if user.Id == 0 {
		return ErrorUserNotFound
	}

	response.Code = http.StatusOK
	response.Message = "success"
	response.Data = &proto.User{
		Id:      user.Id,
		Name:    user.Name,
		Account: user.Account,
		Avatar:  user.Avatar,
	}
	return nil
}

func (h *Handler) UpdateUserInfo(ctx context.Context, request *proto.Request, response *proto.Response) error {

	if request.Name == "" || request.Password == "" || request.Id == 0 {
		return ErrorBadRequest
	}

	user, err := h.UserEntity.GetUserInfo(request.Id, "", "")
	if err != nil {
		return ErrorInternalServerError(err)
	}

	if user.Id == 0 {
		return ErrorUserNotFound
	}

	user, err = h.UserEntity.UpdateUserInfo(request.Id, request.Name, request.Password, request.Avatar)
	if err != nil {
		return ErrorInternalServerError(err)
	}
	response.Code = http.StatusOK
	response.Message = "success"
	response.Data = &proto.User{
		Id:       user.Id,
		Name:     user.Name,
		Password: user.Password,
		Avatar:   user.Avatar,
	}
	return nil
}

func (h *Handler) CheckUserPwd(ctx context.Context, request *proto.Request, response *proto.Response) error {

	if request.Account == "" || request.Password == "" {
		return ErrorBadRequest
	}

	user, err := h.UserEntity.GetUserInfo(0, "", request.Account)
	if err != nil {
		return ErrorInternalServerError(err)
	}

	if user.Id == 0 {
		return ErrorUserNotFound
	}

	if user.Password != request.Password {
		return ErrorUserPasswordError
	}

	response.Code = http.StatusOK
	response.Message = "success"
	response.Data = &proto.User{
		Id:       user.Id,
		Name:     user.Name,
		Password: user.Password,
		Avatar:   user.Avatar,
	}

	return nil
}
