package kafka

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
)

var (
	client      sarama.SyncProducer
	logDataChan chan *logData
)

type logData struct {
	topic string
	data  string
}

// Init ...
func Init(addrs []string, maxSize int) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	client, err = sarama.NewSyncProducer(addrs, config)
	logDataChan = make(chan *logData, maxSize)
	go sendKafkaMsg()
	return
}

// SendChan ...
func SendChan(topic, data string) {
	ld := &logData{
		topic: topic,
		data:  data,
	}
	logDataChan <- ld
}

// sendKafkaMsg ...
func sendKafkaMsg() {
	for {
		select {
		case ld := <-logDataChan:
			msg := &sarama.ProducerMessage{}
			msg.Topic = ld.topic
			msg.Value = sarama.StringEncoder(ld.data)
			pid, offset, err := client.SendMessage(msg)
			if err != nil {
				fmt.Println("send msg err:", err)
				return
			}
			fmt.Println("pid:", pid)
			fmt.Println("offset:", offset)
		default:
			time.Sleep(time.Millisecond * 50)
		}
	}

}
