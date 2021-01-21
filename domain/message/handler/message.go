package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"liaotian/domain/message/proto"
	"liaotian/domain/message/repository"
)

func (h *Handler)Send(ctx context.Context, request *proto.SendRequest, response *proto.Response) error {
	if request.FriendId == 0 || request.SenderId == 0 || request.ReceiverId == 0 || len(request.Content) == 0 {
		return ErrorBadRequest
	}

	//落到kafka   接收人id:
	msgContent, _ := json.Marshal(request)
	msg := &sarama.ProducerMessage{
		Topic: fmt.Sprintf("/message/%v", request.ReceiverId),
		Value: sarama.ByteEncoder(msgContent),
	}

	_, _, err := repository.GetProducer().SendMessage(msg)
	if err != nil {
		return ErrorInternalServerError(err)
	}
	response.Ok = true
	response.Message = "success"

	return nil
}

func (h *Handler)Sub(ctx context.Context, request *proto.SubRequest, response *proto.Response) error {

	if request.UserId == 0 {
		return ErrorBadRequest
	}

	onlineUsersKey := fmt.Sprintf("/online_users/%v/%v", request.UserId%10, request.UserId)

	_, err := repository.GetKv().Put(ctx, onlineUsersKey, string(request.UserId))
	if err != nil{
		return ErrorInternalServerError(err)
	}
	response.Ok = true
	response.Message = "success"
	return nil
}

func (h *Handler)UnSub(ctx context.Context, request *proto.UnSubRequest, response *proto.Response) error {

	if request.UserId == 0 {
		return ErrorBadRequest
	}

	onlineUsersKey := fmt.Sprintf("/online_users/%v/%v", request.UserId%10, request.UserId)

	_, err := repository.GetKv().Delete(ctx, onlineUsersKey)
	if err != nil{
		return ErrorInternalServerError(err)
	}
	response.Ok = true
	response.Message = "success"
	return nil
}

