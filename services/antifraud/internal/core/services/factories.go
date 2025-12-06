package services

import (
	"concurrency-simulator/services/antifraud/utils"
	"database/sql"
)

func NewAntifraudAnalisysService(driver *sql.DB) *AntifraudAnalisysService {
	return &AntifraudAnalisysService{
		log:    utils.NewRequestLogger(),
		driver: driver,
	}
}
