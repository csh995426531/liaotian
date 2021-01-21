package repository

import (
	"errors"
	"github.com/Shopify/sarama"
	"github.com/coreos/etcd/clientv3"
	"time"
)

var (
	ProducerPond []*sarama.SyncProducer
	ConsumerPond []*sarama.Consumer
	EtcdClient *clientv3.Client
)

//初始化
func Init () error {
	startTime := time.Now().Unix()
	for len(ProducerPond) < 10 {
		if producer, err := NewProducer(); err == nil {
			ProducerPond = append(ProducerPond, &producer)
		}
		if time.Now().Unix() - startTime > 10 {
			return errors.New("init kafka producerPond 失败")
		}
	}

	startTime = time.Now().Unix()
	for len(ConsumerPond) < 10 {
		if consumer, err := NewConsumer(); err == nil {
			ConsumerPond = append(ConsumerPond, &consumer)
		}
		if time.Now().Unix() - startTime > 10 {
			return errors.New("init kafka consumerPond 失败")
		}
	}

	client, err := NewClient()
	if err != nil {
		return errors.New("init etcd client 失败")
	}
	EtcdClient = client
	return nil
}

func Close() {
	for i := 0; i < len(ProducerPond); i ++ {
		_= (*ProducerPond[i]).Close()
	}
	ProducerPond = nil
	for i := 0; i < len(ConsumerPond); i ++ {
		_= (*ConsumerPond[i]).Close()
	}
	ConsumerPond = nil
	_= EtcdClient.Close()
}