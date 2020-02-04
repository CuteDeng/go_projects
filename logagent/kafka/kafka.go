package kafka

import (
	"fmt"

	"github.com/Shopify/sarama"
)

var (
	client sarama.SyncProducer
)

// Init ...
func Init(addrs []string) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	client, err = sarama.NewSyncProducer(addrs, config)
	return
}

// SendKafkaMsg ...
func SendKafkaMsg(topic, data string) {
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg err:", err)
		return
	}
	fmt.Println("pid:", pid)
	fmt.Println("offset:", offset)
}
