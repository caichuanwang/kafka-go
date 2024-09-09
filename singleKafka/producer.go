package singlekafka

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/IBM/sarama"
	"github.com/caichuanwang/kafka-go/common"
)

func Producer() {
	client := common.NewClient()

	producer, err := sarama.NewSyncProducerFromClient(client)
	if err != nil {
		log.Fatal(err)
	}
	defer producer.Close()

	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*50)
	defer cancelFunc()

	t := time.NewTicker(time.Second)

	for {
		select {
		case <-timeout.Done():
			fmt.Println("timeout done")
			t.Stop()
			return
		case <-t.C:
			fmt.Println("定时器触发啦！")

			message := sarama.ProducerMessage{
				Topic:     common.Topic,
				Partition: int32(rand.Intn(3)),
				Value:     sarama.StringEncoder(fmt.Sprintf(time.Now().Format(time.DateTime))),
			}

			partition, offset, err := producer.SendMessage(&message)
			if err != nil {
				log.Fatalf("send message err: %+v", err.Error())
			}

			log.Printf("partition: %+v; offset: %+v", partition, offset)
		}
	}

}
