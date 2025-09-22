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
	consumer         *kafka.Consumer
	antifraudService *services.AntifraudAnalisysService
}

func (pc *PaymentConsumer) ProcessMessage(msg *kafka.Message) {
	log.Printf("Received message from topic %s: %s", *msg.TopicPartition.Topic, string(msg.Value))

	var paymentMsg models.PaymentMessage
	if err := json.Unmarshal(msg.Value, &paymentMsg); err != nil {
		log.Printf("Error unmarshalling message: %v", err)
		return
	}

	antifraudData := services.AntifraudAnalisysServiceData{
		FirstName:    paymentMsg.FirstName,
		LastName:     paymentMsg.LastName,
		Amount:       paymentMsg.Amount,
		Installments: paymentMsg.Installments,
		Status:       "pending",
	}

	approved, err := pc.antifraudService.Execute(antifraudData)
	if err != nil {
		log.Printf("Antifraud analysis failed: %v", err)
		pc.logTransactionResult(paymentMsg, false)
		return
	}

	pc.logTransactionResult(paymentMsg, approved)
}

func (pc *PaymentConsumer) logTransactionResult(payment models.PaymentMessage, approved bool) {
	status := "APPROVED"
	if !approved {
		status = "REJECTED"
	}

	logMsg := fmt.Sprintf("ANTIFRAUD_RESULT: Email=%s, Amount=%.2f, Installments=%d, Status=%s",
		payment.Email, payment.Amount, payment.Installments, status)

	log.Println(logMsg)
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

		pc.ProcessMessage(msg)
	}
}

func (pc *PaymentConsumer) Close() {
	if pc.consumer != nil {
		pc.consumer.Close()
	}
}

func NewPaymentConsumer() (*PaymentConsumer, error) {
	logger := utils.NewLogger()

	server := os.Getenv("KAFKA_BROKER")
	groupID := os.Getenv("KAFKA_GROUP_ID")

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

	return &PaymentConsumer{
		consumer:         consumer,
		antifraudService: services.NewAntifraudAnalisysService(),
	}, nil
}
