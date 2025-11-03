package services

import (
	"concurrency-simulator/monorepo/antifraud/utils"
	"database/sql"
)

func NewAntifraudAnalisysService(driver *sql.DB) *AntifraudAnalisysService {
	return &AntifraudAnalisysService{
		log:    utils.NewRequestLogger(),
		driver: driver,
	}
}
