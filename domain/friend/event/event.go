package event

import (
	"github.com/micro/go-micro/broker"
	"liaotian/middlewares/logger/zap"
	"sync"
)

var (
	m        sync.Mutex
	Instance *Event
)

type Event struct {
	PubSub broker.Broker
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
	}
	return
}
