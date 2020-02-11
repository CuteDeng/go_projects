package kafka

import (
	"fmt"
	"go_projects/logtransfer/es"

	"github.com/Shopify/sarama"
)

//
//type LogData struct {
//	Data string `json:"data"`
//}

// Init ...
func Init(addrs []string, topic string) error {
	consumer, err := sarama.NewConsumer(addrs, nil)
	if err != nil {
		fmt.Println("fail to connect kafka consumer", err)
		return err
	}
	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		fmt.Println("fail to get list of partition,err:", err)
		return err
	}
	fmt.Println("partition list", partitionList)
	for partition := range partitionList {
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("fail to start consumer for partition %d,err:%v", partition, err)
			return err
		}
		//defer pc.AsyncClose()
		go func(partitionConsumer sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("partition %d, offset %d, key %v, value %v \n", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
				//ld := map[string]interface{}{
				//	"data": string(msg.Value),
				//}
				ld := es.LogData{
					Topic: topic,
					Data:  string(msg.Value),
				}
				es.SendToChan(&ld)
			}
		}(pc)
	}
	return err
}
