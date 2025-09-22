package main

import (
	"concurrency-simulator/monorepo/antifraud/handler"
	"concurrency-simulator/monorepo/antifraud/utils"
	"sync"
)

func main() {
	logger := utils.NewLogger()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		consumer, err := handler.NewPaymentConsumer()

		if err != nil {
			logger.Fatalf("Failed to create Kafka consumer: %v", err)
		}

		defer consumer.Close()

		logger.Println("Starting Kafka consumer...")
		if err := consumer.Start(); err != nil {
			logger.Fatalf("Failed to start Kafka consumer: %v", err)
		}
	}()

	wg.Wait()
}
