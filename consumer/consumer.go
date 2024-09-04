package consumer

import (
	"log"

	"github.com/IBM/sarama"
	"github.com/caichuanwang/kafka-go/common"
)

func Consumer() {
	client := common.NewClient()

	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		log.Fatalf("err : %+v", err)
	}
	defer consumer.Close()

	pc, err := consumer.ConsumePartition(common.Topic, 0, 0)
	if err != nil {
		log.Fatalf("err : %+v", err)
	}

	go func() {
		for message := range pc.Messages() {
			log.Printf("get one message; topic: %+v, offset: %+v, partition: %+v, value: %s", message.Topic, message.Offset, message.Partition, message.Value)
		}
	}()

	for {

	}
}
