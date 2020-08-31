package handler

import (
	"context"
	proto "liaotian/user-service/proto/user"
)

func (h *Handler) Create(ctx context.Context, request *proto.CreateRequest, response *proto.Response) error {

	user, err := h.repo.Create(request.Name, request.Password)
	if err != nil {
		return err
	}

	response.Code = 200
	response.Message = "SUCCESS"
	response.User = &proto.User{
		Name: user.Username,
	}
	return nil
}

func (h *Handler) Get(ctx context.Context, request *proto.Request, response *proto.Response) error {

	user, err := h.repo.Get(request.Name, request.Password, request.Id)
	if err != nil {
		return err
	}

	response.Code = 200
	response.Message = "SUCCESS"
	response.User = &proto.User{
		Id:       user.ID,
		Name:     user.Username,
		Password: user.Password,
	}
	return nil
}
