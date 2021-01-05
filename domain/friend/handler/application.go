package handler

import (
	"context"
	"liaotian/domain/friend/proto"
)

func (h *Handler) CreateApplicationInfo(ctx context.Context, request *proto.CreateApplicationRequest, response *proto.ApplicationResponse) error {

	return nil
}

func (h *Handler) GetApplicationInfo(ctx context.Context, request *proto.GetApplicationRequest, response *proto.ApplicationResponse) error {

	return nil
}

func (h *Handler) PassApplicationInfo(ctx context.Context, request *proto.PassApplicationInfoRequest, response *proto.Response) error {

	return nil
}

func (h *Handler) RejectApplicationInfo(ctx context.Context, request *proto.RejectApplicationInfoRequest, response *proto.Response) error {

	return nil
}

func (h *Handler) GetApplicationList(ctx context.Context, request *proto.GetApplicationListRequest, response *proto.GetApplicationListResponse) error {

	return nil
}

func (h *Handler) GetFriendList(ctx context.Context, request *proto.GetFriendListRequest, response *proto.GetFriendListResponse) error {

	return nil
}

func (h *Handler) DeleteFriendInfo(ctx context.Context, request *proto.DeleteFriendInfoRequest, response *proto.Response) error {

	return nil
}

func (h *Handler) CreateApplicationSay(ctx context.Context, request *proto.CreateApplicationSayRequest, response *proto.CreateApplicationSayResponse) error {

	return nil
}
