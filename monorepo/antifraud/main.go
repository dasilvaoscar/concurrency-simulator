package main

import (
	"concurrency-simulator/monorepo/antifraud/controllers"
	"concurrency-simulator/monorepo/antifraud/utils"
	"concurrency-simulator/monorepo/shared"
	"log"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	logger := utils.NewLogger()

	var wg sync.WaitGroup
	wg.Add(1)

	go execution(&wg, logger)

	wg.Wait()
}

func execution(wg *sync.WaitGroup, logger *log.Logger) {
	defer wg.Done()

	consumer, err := kafka.NewConsumer(utils.GetKafkaConfig())

	controller := controllers.NewAntifraudController()

	if err != nil {
		logger.Fatalf("Consumer creation error: %s", err)
		panic(err)
	}

	err = consumer.SubscribeTopics([]string{shared.PaymentTopic}, nil)

	if err != nil {
		log.Fatalf("Failed to subscribe to topics: %s", err)
		panic(err)
	}

	log.Printf("Consumer started, listening to topic: %s", shared.PaymentTopic)

	defer consumer.Close()

	for {
		msg, err := consumer.ReadMessage(-1)

		if err != nil {
			log.Printf("Consumer error: %v", err)
			continue
		}

		controller.Execute(msg)
	}
}
