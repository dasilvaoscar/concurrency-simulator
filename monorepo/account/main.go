package main

import (
	"concurrency-simulator/monorepo/account/utils"
	"concurrency-simulator/monorepo/account/controllers"
	"concurrency-simulator/monorepo/shared"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"go.uber.org/zap"
)

func main() {
	logger := utils.NewRequestLogger()

	var wg sync.WaitGroup
	wg.Add(1)

	go execution(&wg, logger)

	wg.Wait()
}

func execution(wg *sync.WaitGroup, logger *zap.Logger) {
	defer wg.Done()

	consumer, err := kafka.NewConsumer(utils.GetKafkaConfig())

	controller := controllers.NewAccountController()

	if err != nil {
		logger.Error("Consumer creation error", zap.Error(err))
		panic(err)
	}

	err = consumer.SubscribeTopics([]string{shared.PaymentTopic}, nil)

	if err != nil {
		logger.Error("Failed to subscribe to topics", zap.Error(err))
		panic(err)
	}

	logger.Error("Consumer started, listening to topic", zap.String("topic", shared.PaymentTopic))

	defer consumer.Close()

	for {
		msg, err := consumer.ReadMessage(-1)

		if err != nil {
			logger.Error("Consumer error", zap.Error(err))
			continue
		}

		controller.ProcessMessage(msg)
	}
}
