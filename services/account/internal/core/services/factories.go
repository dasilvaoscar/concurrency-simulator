package services

import (
	"concurrency-simulator/services/account/utils"
	"database/sql"
)

func NewAccountService(driver *sql.DB) *AccountService {
	return &AccountService{
		log:    utils.NewRequestLogger(),
		driver: driver,
	}
}
