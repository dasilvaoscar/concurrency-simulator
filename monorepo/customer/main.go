package main

import (
	"fmt"
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	server := os.Getenv("KAFKA_BROKER")
	groupId := os.Getenv("KAFKA_GROUP_ID")

	log.Printf("Trying to connect on broker %s with group id %s", server, groupId)

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": server,
		"group.id": groupId,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		log.Fatalf("Consumer creation error: %s", err)
		panic(err)
	}

	defer consumer.Close()

	topic := "payment_topic"
	consumer.SubscribeTopics([]string{topic}, nil)

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message received: %+v\n", msg.Value)
		} else {
			fmt.Printf("Error: %v\n", err)
		}
	}
}
