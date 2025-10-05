package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"concurrency-simulator/monorepo/antifraud/internal/core/models"
	"concurrency-simulator/monorepo/antifraud/internal/core/services"
	"concurrency-simulator/monorepo/shared"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type AntifraudController struct {
	service *services.AntifraudAnalisysService
}

func (pc *AntifraudController) Execute(msg *kafka.Message) {
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

func NewAntifraudController() *AntifraudController {
	dbUrl := os.Getenv("DB_URL")

	db := shared.NewPostgresSingleton(dbUrl)

	return &AntifraudController{
		service: services.NewAntifraudAnalisysService(db),
	}
}
