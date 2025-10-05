package utils

import (
	"concurrency-simulator/monorepo/shared"

	"go.uber.org/zap"
)

func NewRequestLogger() *zap.Logger {
	var loggerInstance = shared.NewLogger("core-svc")

	logger := loggerInstance.With(
		zap.String("topic", shared.PaymentTopic),
	)

	return logger
}
