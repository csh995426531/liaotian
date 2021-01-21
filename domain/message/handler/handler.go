package handler

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/coreos/etcd/clientv3"
	"liaotian/domain/message/event"
	"liaotian/domain/message/repository"
	"liaotian/middlewares/logger/zap"
	"strconv"
)

type Handler struct {

}

func Init() (handler *Handler) {
	handler = new(Handler)
	if err := repository.Init(); err != nil {
		panic(err)
	}
	for i := 0; i < len(repository.ConsumerPond); i ++ {
		go ReadWorker(*repository.ConsumerPond[i], i)
	}
	return
}

func ReadWorker(consumer sarama.Consumer, block int) {
	kv := repository.GetKv()
	for {
		onlineUsersKey := fmt.Sprintf("/online_users/%v/", block)
		resp, err := kv.Get(context.Background(), onlineUsersKey, clientv3.WithPrefix())
		if err != nil {
			zap.SugarLogger.Errorf("repository.GetKv().Get error:%v", err)
			panic(err)
		}

		for _, userInfo := range resp.Kvs {
			uid, _ := strconv.ParseInt(string(userInfo.Value), 10, 64)
			topic := fmt.Sprintf("/message/%v", uid)
			partitionsList, err := consumer.Partitions(topic)
			if err != nil {
				zap.SugarLogger.Errorf("consumer.Partitions error:%v", err)
				panic(err)
			}
			for _, partition := range partitionsList {
				pc, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
				if err != nil {
					zap.SugarLogger.Errorf("consumer.Partitions error:%v", err)
					panic(err)
				}
				go func(pc sarama.PartitionConsumer) {
					for message := range pc.Messages() {
						err := event.Instance.SendMessage(uid, message.Value)
						if err != nil {

						}
					}
				}(pc)
			}
		}
	}
}



