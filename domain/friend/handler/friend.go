package handler

import (
	"context"
	"liaotian/domain/friend/proto"
)

func (h *Handler) GetFriendList(ctx context.Context, request *proto.GetFriendListRequest, response *proto.GetFriendListResponse) error {
	if request.UserId == 0 {
		return ErrorBadRequest
	}

	list, err := h.FriendEntity.GetFriendList(request.UserId)
	if err != nil {
		return ErrorInternalServerError(err)
	}

	response.Message = "success"
	for _, friend := range list {
		if friend.UserIdA == request.UserId {
			response.Data = append(response.Data, friend.UserIdB)
		} else {
			response.Data = append(response.Data, friend.UserIdA)
		}
	}

	return nil
}

func (h *Handler) DeleteFriendInfo(ctx context.Context, request *proto.DeleteFriendInfoRequest, response *proto.Response) error {
	if request.Id == 0 {
		return ErrorBadRequest
	}

	ok, err := h.FriendEntity.DeleteFriendInfo(request.Id)
	if err != nil {
		return ErrorInternalServerError(err)
	}
	response.Message = "success"
	response.Ok = ok
	return nil
}
