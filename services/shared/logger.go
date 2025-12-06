// services/shared/logger.go - COM ZAP
package shared

import (
	"sync"

	"go.uber.org/zap"
)

var (
	logger *zap.Logger
	once   sync.Once
)

func NewLogger(service string) *zap.Logger {
	once.Do(func() {
		var err error
		logger, err = zap.NewProduction()

		if err != nil {
			panic(err)
		}
	})

	return logger.With(zap.String("service", service))
}
