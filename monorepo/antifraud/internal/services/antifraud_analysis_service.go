package services

import (
	"errors"
	"regexp"
)

type AntifraudAnalisysServiceData struct {
	FirstName    string
	LastName     string
	Amount       float64
	Installments int
	Type         string
	Status       string
}

type AntifraudAnalisysService struct{}

func (as *AntifraudAnalisysService) Execute(data AntifraudAnalisysServiceData) (bool, error) {
	highRiskAmount := 50000.0  
	suspiciousInstallmentAmount := 10000.0
	suspiciousInstallments := 6

	// Convert to strategy on the future
	if data.Amount > highRiskAmount {
		return false, errors.New("transação suspeita: valor muito alto, risco elevado")
	}

	if data.Amount > suspiciousInstallmentAmount && data.Installments > suspiciousInstallments {
		return false, errors.New("transação suspeita: valor alto com número de parcelas elevado")
	}

	if as.isSuspiciousName(data.FirstName) || as.isSuspiciousName(data.LastName) {
		return false, errors.New("transação suspeita: nome ou sobrenome incomum")
	}

	return true, nil
}

func (as *AntifraudAnalisysService) isSuspiciousName(name string) bool {
	match, _ := regexp.MatchString(`^[A-Z]$`, name)
	return match
}

func NewAntifraudAnalisysService() *AntifraudAnalisysService {
	return &AntifraudAnalisysService{}
}
