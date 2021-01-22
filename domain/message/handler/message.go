package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"liaotian/domain/message/event"
	"liaotian/domain/message/proto"
	"liaotian/domain/message/repository"
	"liaotian/middlewares/logger/zap"
	"liaotian/middlewares/tool"
)

func (h *Handler)Send(ctx context.Context, request *proto.SendRequest, response *proto.Response) error {
	if request.FriendId == 0 || request.SenderId == 0 || request.ReceiverId == 0 || len(request.Content) == 0 {
		return ErrorBadRequest
	}

	//落到kafka   接收人id:
	msgContent, _ := json.Marshal(request)
	msg := &sarama.ProducerMessage{
		Topic: fmt.Sprintf("message_%v", request.ReceiverId),
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

	// 写入Etcd在线用户
	onlineUsersKey := fmt.Sprintf("/online_users/%v", request.UserId)
	if _, err := repository.GetKv().Put(ctx, onlineUsersKey, string(request.UserId)); err != nil{
		return ErrorInternalServerError(err)
	}

	// 创建读取消息协程
	if err := ReadWorker(request.UserId); err != nil{
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

	onlineUsersKey := fmt.Sprintf("/online_users/%v", request.UserId)

	_, err := repository.GetKv().Delete(ctx, onlineUsersKey)
	if err != nil{
		return ErrorInternalServerError(err)
	}
	response.Ok = true
	response.Message = "success"
	return nil
}

func ReadWorker(UserId int64) error {
	ConsumerGroup, err := repository.NewConsumerGroup(tool.Int64ToInt(UserId))
	if err != nil {
		return ErrorInternalServerError(err)
	}

	go func(UserId int64, ConsumerGroup sarama.ConsumerGroup) {
		kv := repository.GetKv()
		for {
			onlineUsersKey := fmt.Sprintf("/online_users/%v", UserId)
			resp, err := kv.Get(context.Background(), onlineUsersKey)
			if err != nil {
				zap.SugarLogger.Errorf("repository.GetKv().Get error:%v", err)
				panic(err)
			}
			if resp.Count < 1 {
				zap.SugarLogger.Infof("用户%v已离线", UserId)
				panic(err)
			}

			topic := []string{fmt.Sprintf("message_%v", UserId)}
			consumerGroupHandler := consumerGroupHandler{UserId}
			if err = ConsumerGroup.Consume(context.Background(), topic, consumerGroupHandler); err != nil {
				zap.SugarLogger.Infof("ConsumerGroup.Consume 消费失败:%v", err)
				panic(err)
			}
		}
	}(UserId, ConsumerGroup)

	return nil
}

type consumerGroupHandler struct{
	UserId int64
}
func (consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		if err := event.Instance.SendMessage(h.UserId, message.Value); err != nil {
			return err
		}
		// 手动确认消息
		sess.MarkMessage(message, "")
	}
	return nil
}
