package event

import (
	"encoding/json"
	"github.com/micro/go-micro/broker"
	"liaotian/domain/friend/entity"
)

const PassApplication = "PassApplication"

func (e *Event) PassApplication(friend *entity.Friend) (err error) {
	msg := &broker.Message{
		Header: map[string]string{
			"id": string(friend.Id),
		},
	}

	msg.Body, _ = json.Marshal(friend)
	//err = e.PubSub.Publish(PassApplication, msg)
	return
}
