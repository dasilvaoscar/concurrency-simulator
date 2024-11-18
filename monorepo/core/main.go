package main

import (
	"concurrency-simulator/monorepo/core/kafka_producer"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	serverPort := os.Getenv("HTTP_SERVER_PORT")
	serviceName := os.Getenv("SERVICE_NAME")

	fmt.Printf("[%s]: is running on port %s\n", strings.ToUpper(serviceName), serverPort)

	producer, _ := kafka_producer.NewKafkaProducer()
	defer producer.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		deliveryChan := make(chan kafka.Event)

		topic := "payment_topic"

		message := &kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny,
			},
			Value: []byte(fmt.Sprintf("Mensagem")),
		}

		producer.Produce(message, deliveryChan)

		go func() {
			for e := range deliveryChan {
				msg := e.(*kafka.Message)
				if msg.TopicPartition.Error != nil {
					fmt.Printf("Erro na entrega da mensagem: %v\n", msg.TopicPartition.Error)
				} else {
					fmt.Printf("Mensagem enviada com sucesso para %v\n", msg.TopicPartition)
				}
			}
		}()

		producer.Flush(5000)

		close(deliveryChan)
		fmt.Println("Produtor encerrado.")

		w.Write([]byte("This is the " + serviceName + " service"))
	})

	err := http.ListenAndServe(":" + serverPort, nil)

	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
