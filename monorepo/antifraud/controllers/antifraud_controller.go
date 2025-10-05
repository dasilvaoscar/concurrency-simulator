package controllers

import (
	"encoding/json"
	"fmt"
	"os"

	"concurrency-simulator/monorepo/antifraud/internal/core/models"
	"concurrency-simulator/monorepo/antifraud/internal/core/services"
	"concurrency-simulator/monorepo/antifraud/utils"
	"concurrency-simulator/monorepo/shared"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"go.uber.org/zap"
)

type AntifraudController struct {
	logger  *zap.Logger
	service *services.AntifraudAnalisysService
}

func (pc *AntifraudController) ProcessMessage(msg *kafka.Message) {
	pc.logger.Info("Received message from topic", zap.String("topic", *msg.TopicPartition.Topic), zap.String("message", string(msg.Value)))

	var paymentMsg models.Payment
	if err := json.Unmarshal(msg.Value, &paymentMsg); err != nil {
		pc.logger.Error("Error unmarshalling message", zap.Error(err))
		return
	}

	payment := pc.service.Execute(paymentMsg)

	message := fmt.Sprintf("ANTIFRAUD_RESULT: Email=%s, Amount=%.2f, Installments=%d, Status=%s",
		payment.Email, payment.Amount, payment.Installments, *payment.Status)

	pc.logger.Info(message)
}

func NewAntifraudController() *AntifraudController {
	dbUrl := os.Getenv("DB_URL")

	db := shared.NewPostgresSingleton(dbUrl)

	return &AntifraudController{
		logger:  utils.NewRequestLogger(),
		service: services.NewAntifraudAnalisysService(db),
	}
}
