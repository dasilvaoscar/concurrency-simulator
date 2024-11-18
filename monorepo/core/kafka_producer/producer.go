package kafka_producer

import (
	"errors"
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)


func NewKafkaProducer() (*kafka.Producer, error) {
	server := os.Getenv("KAFKA_BROKER")

	log.Printf("Trying to connect on broker %s", server)

	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": server,
	})

	if err != nil {
		log.Fatalf("Producer creation error: %s", err)
		return nil, errors.New("Producer creation error")
	}

	return producer, nil
}
