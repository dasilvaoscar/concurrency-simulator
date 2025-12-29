package main

import (
	"concurrency-simulator/services/antifraud/controllers"
	"concurrency-simulator/services/antifraud/utils"
	"concurrency-simulator/services/shared"
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

	consumer := createConsumer(logger)

	controller := controllers.NewAntifraudController()

	assingPartitions(consumer, logger)
	subscribeToTopic(consumer, logger)

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

func createConsumer(logger *zap.Logger) *kafka.Consumer {
	consumer, err := kafka.NewConsumer(utils.GetKafkaConfig())

	if err != nil {
		logger.Error("Consumer creation error", zap.Error(err))
		panic(err)
	}

	return consumer
}

func assingPartitions(consumer *kafka.Consumer, logger *zap.Logger) {
	topic := shared.PaymentTopic
	err := consumer.Assign([]kafka.TopicPartition{
		{
			Topic:     &topic,
			Partition: shared.PartitionAlias["starting"],
		},
	})

	if err != nil {
		logger.Error("Failed to assign partitions", zap.Error(err))
		panic(err)
	}
}

func subscribeToTopic(consumer *kafka.Consumer, logger *zap.Logger) {
	err := consumer.SubscribeTopics([]string{shared.PaymentTopic}, nil)

	if err != nil {
		logger.Error("Failed to subscribe to topics", zap.Error(err))
		panic(err)
	}
}
