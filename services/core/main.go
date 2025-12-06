package main

import (
	"concurrency-simulator/services/core/handlers"
	"concurrency-simulator/services/core/kafka_producer"
	"concurrency-simulator/services/shared"
	"net/http"
	"os"
	"strings"

	"go.uber.org/zap"
)

func main() {
	logger := shared.NewLogger("core-svc")

	logger.Info("Service is running on port", zap.String("service", strings.ToUpper(os.Getenv("SERVICE_NAME"))), zap.String("port", os.Getenv("HTTP_SERVER_PORT")))

	kafkaProducer, err := kafka_producer.NewKafkaProducer()

	if err != nil {
		logger.Error("Error creating producer", zap.Error(err))
		os.Exit(1)
	}

	defer kafkaProducer.Close()

	http.HandleFunc("/payment", func(w http.ResponseWriter, r *http.Request) {
		handlers.PaymentHandler(w, r, kafkaProducer)
	})

	err = http.ListenAndServe(":"+os.Getenv("HTTP_SERVER_PORT"), nil)

	if err != nil {
		logger.Error("Error starting server", zap.Error(err))
	}
}
