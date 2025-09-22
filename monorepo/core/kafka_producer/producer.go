package kafka_producer

import (
	"concurrency-simulator/monorepo/core/utils"
	"errors"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func NewKafkaProducer() (*kafka.Producer, error) {
	server := os.Getenv("KAFKA_BROKER")

	logger := utils.NewLogger()

	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":  server,
		"enable.idempotence": "true",
		"acks":               "all",
		"linger.ms":          0,
		"batch.size":         1,
	})

	if err != nil {
		logger.Fatalf("Producer creation error: %s", err)
		return nil, errors.New("Producer creation error")
	}

	return producer, nil
}
