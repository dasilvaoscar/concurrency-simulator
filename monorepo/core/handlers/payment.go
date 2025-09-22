package handlers

import (
	"concurrency-simulator/monorepo/core/utils"
	"concurrency-simulator/monorepo/shared"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var logger = utils.NewLogger()

func PaymentHandler(w http.ResponseWriter, r *http.Request, kafkaProducer *kafka.Producer) {

	topic := shared.PaymentTopic

	bodyMap, err := parseBody(r)

	if err != nil {
		logger.Printf("Error parsing body: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := validateBody(bodyMap); err != nil {
		logger.Printf("Error: amount, installments and email are required")
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
		logger.Printf("Error: amount is not a float64")
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
		logger.Printf("Error marshalling message: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: jsonMessage,
	}

	deliveryChan := make(chan kafka.Event)
	kafkaProducer.Produce(message, deliveryChan)

	go func() {
		for e := range deliveryChan {
			msg := e.(*kafka.Message)
			if msg.TopicPartition.Error != nil {
				logger.Printf("Erro na entrega da mensagem: %v\n", msg.TopicPartition.Error)
			}
		}
	}()

	pendingMessages := kafkaProducer.Flush(1000)
	if pendingMessages > 0 {
		logger.Printf("AVISO: %d mensagens ainda pendentes após flush\n", pendingMessages)
	}

	close(deliveryChan)
	logger.Println("Produtor encerrado.")

	w.Write([]byte("sent to kafka"))
}

func parseBody(r *http.Request) (map[string]interface{}, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Printf("Error reading body: %v\n", err)
		return nil, err
	}

	defer r.Body.Close()

	var bodyMap map[string]interface{}

	err = json.Unmarshal(body, &bodyMap)
	if err != nil {
		logger.Printf("Error unmarshalling body: %v\n", err)
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
