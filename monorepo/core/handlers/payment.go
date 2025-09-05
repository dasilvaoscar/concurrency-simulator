package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func PaymentHandler(w http.ResponseWriter, r *http.Request, kafkaProducer *kafka.Producer) {
	topic := "payment_topic"

	bodyMap, err := parseBody(r)

	if err != nil {
		fmt.Printf("Error parsing body: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := validateBody(bodyMap); err != nil {
		fmt.Printf("Error: amount, installments and email are required")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	amount := bodyMap["amount"]
	installments := bodyMap["installments"]
	email := bodyMap["email"]

	amountFloat, ok := amount.(float64)
	if !ok {
		fmt.Printf("Error: amount is not a float64")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	originalMessage := map[string]interface{}{
		"amount":       amountFloat,
		"installments": installments,
		"email":        email,
	}

	jsonMessage, err := json.Marshal(originalMessage)
	if err != nil {
		fmt.Printf("Error marshalling message: %v\n", err)
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
				fmt.Printf("Erro na entrega da mensagem: %v\n", msg.TopicPartition.Error)
			}
		}
	}()

	pendingMessages := kafkaProducer.Flush(1000)
	if pendingMessages > 0 {
		fmt.Printf("AVISO: %d mensagens ainda pendentes após flush\n", pendingMessages)
	}

	close(deliveryChan)
	fmt.Println("Produtor encerrado.")

	w.Write([]byte("This is the " + os.Getenv("SERVICE_NAME") + " service"))
}

func parseBody(r *http.Request) (map[string]interface{}, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error reading body: %v\n", err)
		return nil, err
	}

	defer r.Body.Close()

	var bodyMap map[string]interface{}

	err = json.Unmarshal(body, &bodyMap)
	if err != nil {
		fmt.Printf("Error unmarshalling body: %v\n", err)
		return nil, err
	}

	return bodyMap, nil
}

func validateBody(bodyMap map[string]interface{}) error {
	if bodyMap["amount"] == nil || bodyMap["installments"] == nil || bodyMap["email"] == nil {
		return errors.New("amount, installments and email are required")
	}

	return nil
}
