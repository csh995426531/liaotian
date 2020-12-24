package handler

import (
	"context"
	"liaotian/domain/user/proto"
)

/**
	用户领域服务
 */

func (h *Handler) CreateUserInfo (ctx context.Context, request *proto.Request, response *proto.Response) error {

	user, err := h.UserEntity.GetUserInfo(0, "", request.Account)
	if err != nil {
		return err
	}
	if user != nil {
		response.Code = 500
		response.Message = "账户已被注册！"
		response.Data = nil
		return nil
	}

	user, err = h.UserEntity.CreateUserInfo(request.Name, request.Account, request.Password, request.Avatar)
	if err != nil {
		return err
	}
	response.Code = 200
	response.Message = "success"
	response.Data = &proto.User{
		Id: user.Id,
		Name: user.Name,
		Password: user.Password,
		Avatar: user.Avatar,
	}

	return nil
}

func (h *Handler) GetUserInfo (ctx context.Context, request *proto.Request, response *proto.Response) error {

	user, err := h.UserEntity.GetUserInfo(request.Id, request.Name, request.Account)
	if err != nil {
		return err
	}

	response.Code = 200
	response.Message = "success"

	if user.Id > 0 {
		response.Data = &proto.User{
			Id: user.Id,
			Name: user.Name,
			Password: user.Password,
			Avatar: user.Avatar,
		}
	}
	return nil
}

func (h *Handler) UpdateUserInfo (ctx context.Context, request *proto.Request, response *proto.Response) error {

	user, err := h.UserEntity.GetUserInfo(request.Id, "", "")
	if err != nil {
		return err
	}

	if user.Id == 0 {
		response.Code = 500
		response.Message = "用户不存在"
		return nil
	}

	user, err = h.UserEntity.UpdateUserInfo(request.Id, request.Name, request.Password, request.Avatar)
	if err != nil {
		return err
	}
	response.Code = 200
	response.Message = "success"
	response.Data = &proto.User{
		Id: user.Id,
		Name: user.Name,
		Password: user.Password,
		Avatar: user.Avatar,
	}
	return nil
}

func (h *Handler) CheckUserPwd (ctx context.Context, request *proto.Request, response *proto.Response) error {

	user, err := h.UserEntity.GetUserInfo(0, "", request.Account)
	if err != nil {
		return err
	}

	if user.Id == 0 {
		response.Code = 500
		response.Message = "用户不存在"
		return nil
	}

	if user.Password != request.Password {
		response.Code = 500
		response.Message = "密码错误"
	} else {
		response.Code = 200
		response.Message = "success"
		response.Data = &proto.User{
			Id: user.Id,
			Name: user.Name,
			Password: user.Password,
			Avatar: user.Avatar,
		}
	}

	return nil
}
