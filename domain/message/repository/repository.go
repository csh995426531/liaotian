package repository

import (
	"errors"
	"github.com/Shopify/sarama"
	"github.com/coreos/etcd/clientv3"
	"liaotian/middlewares/logger/zap"
	"time"
)

var (
	ProducerPond      []*sarama.SyncProducer
	ConsumerGroupPond []*sarama.ConsumerGroup
	EtcdClient        *clientv3.Client
)

//初始化
func Init() error {
	startTime := time.Now().Unix()
	for len(ProducerPond) < 100 {
		if producer, err := NewProducer(); err == nil {
			ProducerPond = append(ProducerPond, &producer)
		}
		if time.Now().Unix()-startTime > 10 {
			return errors.New("init kafka producerPond 失败")
		}
	}

	//startTime = time.Now().Unix()
	//for len(ConsumerGroupPond) < 10 {
	//	if ConsumerGroup, err := NewConsumerGroup(len(ConsumerGroupPond)); err == nil {
	//		ConsumerGroupPond = append(ConsumerGroupPond, &ConsumerGroup)
	//	}
	//	if time.Now().Unix() - startTime > 10 {
	//		return errors.New("init kafka consumerPond 失败")
	//	}
	//}

	client, err := NewClient()
	if err != nil {
		return errors.New("init etcd client 失败")
	}
	EtcdClient = client
	zap.ZapLogger.Info("repository 初始化成功")
	return nil
}

func Close() {
	for i := 0; i < len(ProducerPond); i++ {
		_ = (*ProducerPond[i]).Close()
	}
	ProducerPond = nil
	for i := 0; i < len(ConsumerGroupPond); i++ {
		_ = (*ConsumerGroupPond[i]).Close()
	}
	ConsumerGroupPond = nil
	_ = EtcdClient.Close()
}
