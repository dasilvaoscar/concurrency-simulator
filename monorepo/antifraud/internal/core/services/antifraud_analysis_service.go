package services

import (
	"concurrency-simulator/monorepo/antifraud/utils"
	"errors"
	"log"
	"regexp"
)

type AntifraudAnalisysServiceData struct {
	FirstName    string
	LastName     string
	Email        string
	Amount       float64
	Installments int
	Status       string
}

type AntifraudAnalisysService struct {
	log *log.Logger
}

func (as *AntifraudAnalisysService) Execute(data AntifraudAnalisysServiceData) (bool, error) {
	highRiskAmount := 50000.0
	suspiciousInstallmentAmount := 10000.0
	suspiciousInstallments := 6

	if data.Amount > highRiskAmount {
		as.log.Printf("Transação suspeita: valor muito alto, risco elevado - Email: %s - Amount: %.2f", data.Email, data.Amount)
		return false, errors.New("transação suspeita: valor muito alto, risco elevado")
	}

	if data.Amount > suspiciousInstallmentAmount && data.Installments > suspiciousInstallments {
		as.log.Printf("Transação suspeita: valor alto com número de parcelas elevado - Email: %s - Amount: %.2f - Installments: %d", data.Email, data.Amount, data.Installments)
		return false, errors.New("transação suspeita: valor alto com número de parcelas elevado")
	}

	if as.isSuspiciousName(data.FirstName) || as.isSuspiciousName(data.LastName) {
		as.log.Printf("Transação suspeita: nome ou sobrenome incomum - Email: %s - FirstName: %s - LastName: %s", data.Email, data.FirstName, data.LastName)
		return false, errors.New("transação suspeita: nome ou sobrenome incomum")
	}

	return true, nil
}

func (as *AntifraudAnalisysService) isSuspiciousName(name string) bool {
	match, _ := regexp.MatchString(`^[A-Z]$`, name)
	return match
}

func NewAntifraudAnalisysService() *AntifraudAnalisysService {
	return &AntifraudAnalisysService{
		log: utils.NewLogger(),
	}
}
