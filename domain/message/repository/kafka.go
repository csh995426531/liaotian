package repository

import (
	"github.com/Shopify/sarama"
	"k8s.io/apimachinery/pkg/util/rand"
	"strings"
)

func NewProducer () (producer sarama.SyncProducer, err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	producer, err = sarama.NewSyncProducer([]string{"192.168.66.104:9092"}, config)
	return
}

func NewConsumer() (consumer sarama.Consumer, err error) {
	consumer, err = sarama.NewConsumer(strings.Split("192.168.66.104:9092", ","), nil)
	return
}

func GetProducer() (producer sarama.SyncProducer) {
	random := rand.IntnRange(0, len(ProducerPond))
	producer = *ProducerPond[random]
	return
}

func GetConsumer() (consumer sarama.Consumer) {
	random := rand.IntnRange(0, len(ConsumerPond))
	consumer = *ConsumerPond[random]
	return
}
