package controllers

import (
	"concurrency-simulator/monorepo/antifraud/internal/core/services"
	"concurrency-simulator/monorepo/antifraud/utils"
	"concurrency-simulator/monorepo/shared"
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
