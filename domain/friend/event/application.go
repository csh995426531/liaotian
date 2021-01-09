package event

import (
	"encoding/json"
	"github.com/micro/go-micro/broker"
)

const PassApplication = "PassApplication"

func (e *Event) PassApplication(userIdA, userIdB int64) (err error) {
	msg := &broker.Message{
		Header: map[string]string{
			"user_id_a": string(userIdA),
			"user_id_b": string(userIdB),
		},
	}
	msg.Body, _ = json.Marshal(msg.Header)
	err = e.PubSub.Publish(PassApplication, msg)
	return
}
