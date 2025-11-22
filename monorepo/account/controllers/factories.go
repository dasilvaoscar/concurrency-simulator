package controllers

import (
	"concurrency-simulator/monorepo/account/internal/core/services"
	"concurrency-simulator/monorepo/antifraud/utils"
	"concurrency-simulator/monorepo/shared"
	"os"
)

func NewAccountController() *AccountController {
	dbUrl := os.Getenv("DB_URL")

	db := shared.NewPostgresSingleton(dbUrl)

	return &AccountController{
		logger:  utils.NewRequestLogger(),
		service: services.NewAccountService(db),
	}
}
