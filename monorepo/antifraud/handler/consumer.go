package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"concurrency-simulator/monorepo/antifraud/internal/core/models"
	"concurrency-simulator/monorepo/antifraud/internal/core/services"
	"concurrency-simulator/monorepo/antifraud/utils"
	"concurrency-simulator/monorepo/shared"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type PaymentConsumer struct {
	consumer *kafka.Consumer
	service  *services.AntifraudAnalisysService
}

func (pc *PaymentConsumer) Start() error {
	err := pc.consumer.SubscribeTopics([]string{shared.PaymentTopic}, nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to topics: %s", err)
		return err
	}

	log.Printf("Consumer started, listening to topic: %s", shared.PaymentTopic)

	for {
		msg, err := pc.consumer.ReadMessage(-1)
		if err != nil {
			log.Printf("Consumer error: %v", err)
			continue
		}

		pc.processMessage(msg)
	}
}

func (pc *PaymentConsumer) Close() {
	if pc.consumer != nil {
		pc.consumer.Close()
	}
}

func (pc *PaymentConsumer) processMessage(msg *kafka.Message) {
	log.Printf("Received message from topic %s: %s", *msg.TopicPartition.Topic, string(msg.Value))

	var paymentMsg models.Payment
	if err := json.Unmarshal(msg.Value, &paymentMsg); err != nil {
		log.Printf("Error unmarshalling message: %v", err)
		return
	}

	payment, err := pc.service.Execute(paymentMsg)
	if err != nil {
		log.Printf("Antifraud analysis failed: %v", err)

		message := fmt.Sprintf("ANTIFRAUD_RESULT: Email=%s, Amount=%.2f, Installments=%d, Status=%s",
			payment.Email, payment.Amount, payment.Installments, payment.Status)

		log.Println(message)

		return
	}

	message := fmt.Sprintf("ANTIFRAUD_RESULT: Email=%s, Amount=%.2f, Installments=%d, Status=%s",
		payment.Email, payment.Amount, payment.Installments, payment.Status)

	log.Println(message)
}

func NewPaymentConsumer() (*PaymentConsumer, error) {
	logger := utils.NewLogger()

	server := os.Getenv("KAFKA_BROKER")
	groupID := os.Getenv("KAFKA_GROUP_ID")
	dbUrl := os.Getenv("DB_URL")

	if groupID == "" {
		groupID = "antifraud-group"
	}

	logger.Printf("Trying to connect to broker %s with group ID %s", server, groupID)

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": server,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		logger.Fatalf("Consumer creation error: %s", err)
		return nil, err
	}

	db := shared.NewPostgresSingleton(dbUrl)

	return &PaymentConsumer{
		consumer: consumer,
		service:  services.NewAntifraudAnalisysService(db),
	}, nil
}
