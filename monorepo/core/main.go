package main

import (
	"concurrency-simulator/monorepo/core/handlers"
	"concurrency-simulator/monorepo/core/kafka_producer"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	fmt.Printf("[%s]: is running on port %s\n", strings.ToUpper(os.Getenv("SERVICE_NAME")), os.Getenv("HTTP_SERVER_PORT"))

	kafkaProducer, err := kafka_producer.NewKafkaProducer()

	if err != nil {
		fmt.Printf("Error creating producer: %v\n", err)
		os.Exit(1)
	}

	defer kafkaProducer.Close()
	
	http.HandleFunc("/payment", func(w http.ResponseWriter, r *http.Request) {
		handlers.PaymentHandler(w, r, kafkaProducer)
	})
	
	err = http.ListenAndServe(":"+os.Getenv("HTTP_SERVER_PORT"), nil)

	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
