package common

import (
	"log"
	"strings"

	"github.com/IBM/sarama"
)

const (
	Topic   = "hello"
	Address = "localhost:9092"
	Group   = "group1"
)

func NewClient() sarama.Client {
	address := strings.Split(Address, ",")
	conf := sarama.NewConfig()
	conf.Producer.Return.Successes = true

	client, err := sarama.NewClient(address, conf)
	if err != nil {
		log.Fatalf("new client err: %+v", err)
	}

	return client
}
