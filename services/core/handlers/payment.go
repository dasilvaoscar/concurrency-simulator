package handlers

import (
	"concurrency-simulator/services/core/utils"
	"concurrency-simulator/services/shared"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"go.uber.org/zap"
)

func PaymentHandler(w http.ResponseWriter, r *http.Request, kafkaProducer *kafka.Producer) {
	logger := utils.NewRequestLogger(r)

	topic := shared.PaymentTopic

	bodyMap, err := parseBody(r)

	if err != nil {
		logger.Error("Failed to parse request body",
			zap.Error(err),
			zap.String("content_type", r.Header.Get("Content-Type")),
		)

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := validateBody(bodyMap); err != nil {
		logger.Warn("Invalid request body",
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	amount := bodyMap["amount"]
	installments := bodyMap["installments"]
	email := bodyMap["email"]
	firstName := bodyMap["first_name"]
	lastName := bodyMap["last_name"]

	amountFloat, ok := amount.(float64)
	if !ok {
		logger.Error("Error: amount is not a float64")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	originalMessage := map[string]interface{}{
		"amount":       amountFloat,
		"installments": installments,
		"email":        email,
		"first_name":   firstName,
		"last_name":    lastName,
	}

	jsonMessage, err := json.Marshal(originalMessage)
	if err != nil {
		logger.Error("Failed to marshal payment message",
			zap.Error(err),
			zap.Int("message_size", len(jsonMessage)),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: shared.PartitionAlias["starting"],
		},
		Value: jsonMessage,
	}

	deliveryChan := make(chan kafka.Event)
	kafkaProducer.Produce(message, deliveryChan)

	go func() {
		for e := range deliveryChan {
			msg := e.(*kafka.Message)
			if msg.TopicPartition.Error != nil {
				logger.Error("Kafka delivery failed",
					zap.Error(msg.TopicPartition.Error),
					zap.String("topic", topic),
					zap.Int32("partition", msg.TopicPartition.Partition),
				)
			} else {
				logger.Info("Kafka delivery confirmed",
					zap.String("topic", topic),
					zap.Int32("partition", msg.TopicPartition.Partition),
					zap.Int64("offset", int64(msg.TopicPartition.Offset)),
				)
			}
		}
	}()

	pendingMessages := kafkaProducer.Flush(1000)
	if pendingMessages > 0 {
		logger.Warn("Messages still pending after flush",
			zap.Int("pending_count", pendingMessages),
		)
	}

	close(deliveryChan)
	logger.Info("Payment processing completed",
		zap.String("status", "success"),
	)

	w.Write([]byte("sent to kafka"))
}

func parseBody(r *http.Request) (map[string]interface{}, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	defer r.Body.Close()

	var bodyMap map[string]interface{}

	err = json.Unmarshal(body, &bodyMap)
	if err != nil {
		return nil, err
	}

	return bodyMap, nil
}

func validateBody(bodyMap map[string]interface{}) error {
	if bodyMap["amount"] == nil || bodyMap["installments"] == nil || bodyMap["email"] == nil || bodyMap["first_name"] == nil || bodyMap["last_name"] == nil {
		return errors.New("amount, installments, email, first_name and last_name are required")
	}

	return nil
}
