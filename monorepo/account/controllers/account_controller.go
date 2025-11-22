package controllers

import (
	"encoding/json"
	"fmt"

	"concurrency-simulator/monorepo/account/internal/core/services"
	"concurrency-simulator/monorepo/shared/topic_messages"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"go.uber.org/zap"
)

type AccountController struct {
	logger  *zap.Logger
	service *services.AccountService
}

func (pc *AccountController) ProcessMessage(msg *kafka.Message) {
	var accountMessage topic_messages.Payment
	if err := json.Unmarshal(msg.Value, &accountMessage); err != nil {
		return
	}

	result := pc.service.Execute(accountMessage)

	message := fmt.Sprintf("ACCOUNT_RESULT: %v", result)

	pc.logger.Info(message)
}
