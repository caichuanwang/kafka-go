package main

import (
	"log"
	"time"

	"github.com/caichuanwang/kafka-go/consumer"
	"github.com/caichuanwang/kafka-go/producer"
)

func main() {
	go producer.Producer()

	go consumer.ConsumerGrpInstance()

	log.Print("server is starting")
	time.Sleep(time.Second * 20)
}
