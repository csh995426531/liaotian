package event

import (
	"encoding/json"
	"github.com/docker/distribution/context"
	"github.com/micro/go-micro/broker"
	messageService "liaotian/domain/message/proto"
	"liaotian/middlewares/logger/zap"
)

const PassApplication = "PassApplication"

func (e *Event) PassApplication() {

	type Friend struct {
		Id      int64 `json:"id"`
		UserIdA int64 `json:"user_id_a"`
		UserIdB int64 `json:"user_id_b"`
	}

	for {
		sub, err := e.PubSub.Subscribe(PassApplication, func(event broker.Event) error {

			var friend Friend
			_ = json.Unmarshal(event.Message().Body, &friend)

			res, err := e.DomainMessage.Send(context.Background(), &messageService.SendRequest{
				FriendId:   friend.Id,
				SenderId:   friend.UserIdA,
				ReceiverId: friend.UserIdB,
				Content:    "你们已经成为朋友啦！",
			})
			if err != nil || res.Ok != true {
				zap.SugarLogger.Panicf("订阅事件PassApplication,DomainMessage.Send失败，error: %v", err)
			}
			res, err = e.DomainMessage.Send(context.Background(), &messageService.SendRequest{
				FriendId:   friend.Id,
				SenderId:   friend.UserIdB,
				ReceiverId: friend.UserIdA,
				Content:    "你们已经成为朋友啦！",
			})
			if err != nil || res.Ok != true {
				zap.SugarLogger.Panicf("订阅事件PassApplication,DomainMessage.Send失败，error: %v", err)
			}
			return nil
		})

		if err != nil {
			zap.SugarLogger.Panicf("订阅事件PassApplication失败，error: %v", err)
		}
		_ = sub.Unsubscribe()
	}

}
