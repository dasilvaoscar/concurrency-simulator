package utils

import (
	"log"
	"os"
	"sync"
)

var (
	logger *log.Logger
	once   sync.Once
)

// NewLogger returns a singleton logger instance
func NewLogger() *log.Logger {
	once.Do(func() {
		logger = log.New(os.Stdout, "core-svc: ", log.LstdFlags)
	})
	return logger
}
