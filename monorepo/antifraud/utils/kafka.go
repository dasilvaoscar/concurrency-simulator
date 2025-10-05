package utils

import (
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func GetKafkaConfig() *kafka.ConfigMap {
	server := os.Getenv("KAFKA_BROKER")
	groupID := os.Getenv("KAFKA_GROUP_ID")

	return &kafka.ConfigMap{
		"bootstrap.servers": server,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	}
}
