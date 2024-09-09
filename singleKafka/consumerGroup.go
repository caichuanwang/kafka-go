package singlekafka

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM/sarama"
	"github.com/caichuanwang/kafka-go/common"
)

type ConsumerGrp struct {
}

func (c *ConsumerGrp) Setup(session sarama.ConsumerGroupSession) error {
	fmt.Println("消费启动")
	return nil
}

func (c *ConsumerGrp) Cleanup(session sarama.ConsumerGroupSession) error {
	fmt.Println("消费者关闭")
	return nil
}

func (c *ConsumerGrp) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("接受topic= %+v, partition=%+v, offset=%d, value=%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Value)
	}
	return nil
}

func ConsumerGrpInstance() {
	client := common.NewClient()
	ctx := context.Background()

	g, err := sarama.NewConsumerGroupFromClient(common.Group, client)
	if err != nil {
		log.Fatal(err)
	}
	defer g.Close()

	con := &ConsumerGrp{}
	go func() {
		for {
			if err := g.Consume(ctx, []string{common.Topic}, con); err != nil {
				log.Fatal(err)
			}
		}
	}()

	for {

	}
}
