package services

import (
	"concurrency-simulator/monorepo/antifraud/internal/core/models"
	"concurrency-simulator/monorepo/antifraud/utils"
	"database/sql"
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
}

// TODO: refactor dependencies to use interfaces
type AntifraudAnalisysService struct {
	log    *log.Logger
	driver *sql.DB
}

func (as *AntifraudAnalisysService) Execute(payment models.Payment) (models.Payment, error) {
	highRiskAmount := 50000.0
	suspiciousInstallmentAmount := 10000.0
	suspiciousInstallments := 6

	status := `REJECTED`

	if payment.Amount > highRiskAmount {
		payment.Status = &status
		as.saveDataToDatabase(payment)
		as.log.Printf("Transação suspeita: valor muito alto, risco elevado - Email: %s - Amount: %.2f", payment.Email, payment.Amount)
		return payment, errors.New("transação suspeita: valor muito alto, risco elevado")
	}

	if payment.Amount > suspiciousInstallmentAmount && payment.Installments > suspiciousInstallments {
		payment.Status = &status
		as.saveDataToDatabase(payment)
		as.log.Printf("Transação suspeita: valor alto com número de parcelas elevado - Email: %s - Amount: %.2f - Installments: %d", payment.Email, payment.Amount, payment.Installments)
		return payment, errors.New("transação suspeita: valor alto com número de parcelas elevado")
	}

	if as.isSuspiciousName(payment.FirstName) || as.isSuspiciousName(payment.LastName) {
		payment.Status = &status
		as.saveDataToDatabase(payment)
		as.log.Printf("Transação suspeita: nome ou sobrenome incomum - Email: %s - FirstName: %s - LastName: %s", payment.Email, payment.FirstName, payment.LastName)
		return payment, errors.New("transação suspeita: nome ou sobrenome incomum")
	}

	status = "APPROVED"

	payment.Status = &status
	as.saveDataToDatabase(payment)

	return payment, nil
}

func (as *AntifraudAnalisysService) isSuspiciousName(name string) bool {
	match, _ := regexp.MatchString(`^[A-Z]$`, name)
	return match
}

func (as *AntifraudAnalisysService) saveDataToDatabase(data models.Payment) (bool, error) {
	query := `
		INSERT INTO analysis (first_name, last_name, email, amount, installments, status, created_at) 
		VALUES ($1, $2, $3, $4, $5, $6, NOW())
	`

	_, err := as.driver.Exec(query, data.FirstName, data.LastName, data.Email, data.Amount, data.Installments, data.Status)

	if err != nil {
		as.log.Fatal(err.Error())
		panic("save-data-error")
	}

	return true, nil
}

func NewAntifraudAnalisysService(driver *sql.DB) *AntifraudAnalisysService {
	return &AntifraudAnalisysService{
		log:    utils.NewLogger(),
		driver: driver,
	}
}
