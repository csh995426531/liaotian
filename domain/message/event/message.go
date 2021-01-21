package event

import (
	"fmt"
	"github.com/micro/go-micro/broker"
)

func (e *Event) SendMessage(userId int64, content []byte) (err error) {
	msg := &broker.Message{
		Body: content,
	}
	topic := fmt.Sprintf("/message/%v", userId)
	err = e.PubSub.Publish(topic,msg)
	return
}
