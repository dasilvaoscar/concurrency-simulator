package services

import (
	"concurrency-simulator/monorepo/account/utils"
	"database/sql"
)

func NewAccountService(driver *sql.DB) *AccountService {
	return &AccountService{
		log:    utils.NewRequestLogger(),
		driver: driver,
	}
}
