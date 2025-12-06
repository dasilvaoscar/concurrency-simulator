package controllers

import (
	"concurrency-simulator/services/antifraud/internal/core/services"
	"concurrency-simulator/services/antifraud/utils"
	"concurrency-simulator/services/shared"
	"os"
)

func NewAntifraudController() *AntifraudController {
	dbUrl := os.Getenv("DB_URL")

	db := shared.NewPostgresSingleton(dbUrl)

	return &AntifraudController{
		logger:  utils.NewRequestLogger(),
		service: services.NewAntifraudAnalisysService(db),
	}
}
