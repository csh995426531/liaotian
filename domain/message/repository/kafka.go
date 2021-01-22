package repository

import (
	"github.com/Shopify/sarama"
	"k8s.io/apimachinery/pkg/util/rand"
	"strconv"
)

func NewProducer () (producer sarama.SyncProducer, err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	producer, err = sarama.NewSyncProducer([]string{"192.168.66.104:9092"}, config)
	return
}

func NewConsumerGroup(groupId int) (consumerGroup sarama.ConsumerGroup, err error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = false
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange // 分区分配策略
	config.Version = sarama.V2_3_0_0
	client, err := sarama.NewClient([]string{"192.168.66.104:9092"}, config)
	if err != nil {
		return nil, err
	}

	consumerGroup, err = sarama.NewConsumerGroupFromClient(strconv.Itoa(groupId), client)
	return
}

func GetProducer() (producer sarama.SyncProducer) {
	random := rand.IntnRange(0, len(ProducerPond))
	producer = *ProducerPond[random]
	return
}
