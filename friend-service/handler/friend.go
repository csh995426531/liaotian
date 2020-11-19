package handler

import (
	"context"
	proto "liaotian/friend-service/proto/friend"
)

func (h *Handler) Add(ctx context.Context, request *proto.AddRequest, response *proto.Response) error {
	_, err := h.repo.Add(request.OperatorId, request.BuddyId)
	if err != nil {
		return err
	}

	response.Code = 200
	response.Message = "SUCCESS"
	return nil
}

func (h *Handler) Del(ctx context.Context, request *proto.Request, response *proto.Response) error {

	err := h.repo.Del(request.FriendId)
	if err != nil {
		return err
	}

	response.Code = 200
	response.Message = "SUCCESS"
	return nil
}

func (h *Handler) List(ctx context.Context, request *proto.ListRequest, response *proto.Response) error {

	friends, err := h.repo.List(request.OperatorId, request.Offset, request.Limit)
	if err != nil {
		return err
	}

	response.Code = 200
	response.Message = "SUCCESS"

	for _, friend := range friends {
		response.List = append(response.List, &proto.Friend{Id: friend.ID, UserId: friend.BuddyId})
	}

	return nil
}

func (h *Handler) Get(ctx context.Context, request *proto.Request, response *proto.Response) error {

	friend, err := h.repo.Get(request.FriendId)
	if err != nil {
		return err
	}

	response.Code = 200
	response.Message = "SUCCESS"
	response.Friend = &proto.Friend{
		Id:     friend.ID,
		UserId: friend.BuddyId,
	}

	return nil
}
