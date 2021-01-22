package event

import (
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/client"
	messageService "liaotian/domain/message/proto"
	"liaotian/middlewares/logger/zap"
	"sync"
)

var (
	m        sync.Mutex
	Instance *Event
)

type Event struct {
	PubSub        broker.Broker
	DomainMessage messageService.MessageService
}

func Init(broker broker.Broker) {
	m.Lock()
	defer m.Unlock()
	if Instance != nil {
		zap.ZapLogger.Error("Event已经初始化过了")
		return
	}

	if err := broker.Connect(); err != nil {
		zap.SugarLogger.Errorf("Event broker连接失败，%v", err)
		return
	}

	Instance = &Event{
		broker,
		messageService.NewMessageService("domain.message.service", client.DefaultClient),
	}

	go Instance.PassApplication()

	return
}
