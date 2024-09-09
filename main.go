package main

import (
	"log"
	"time"

	singlekafka "github.com/caichuanwang/kafka-go/singleKafka"
)

func main() {
	// 单个kafka实例
	go singlekafka.Producer()

	go singlekafka.ConsumerGrpInstance()

	log.Print("server is starting")
	time.Sleep(time.Second * 20)
}
