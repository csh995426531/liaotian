package event

import (
	"fmt"
	"github.com/micro/go-micro/broker"
	"liaotian/middlewares/logger/zap"
)

func (e *Event) ReadNewMessage(userId int64) (data []byte) {

	topic := fmt.Sprintf("/message/%v", userId)
	sub, err := e.PubSub.Subscribe(topic, func(event broker.Event) error {
		data = event.Message().Body
		return nil
	})

	if err != nil {
		zap.SugarLogger.Panicf("订阅事件ReadNewMessage失败，error: %v", err)
	}
	_ = sub.Unsubscribe()
	return data
}
