package services

import (
	"concurrency-simulator/monorepo/shared/topic_messages"
	"database/sql"
	"regexp"

	"go.uber.org/zap"
)

type AntifraudAnalisysServiceData struct {
	FirstName    string
	LastName     string
	Email        string
	Amount       float64
	Installments int
}

// TODO: do refactor to use interfaces
type AntifraudAnalisysService struct {
	log    *zap.Logger
	driver *sql.DB
}

func (as *AntifraudAnalisysService) Execute(payment topic_messages.Payment) topic_messages.Payment {
	highRiskAmount := 50000.0
	suspiciousInstallmentAmount := 10000.0
	suspiciousInstallments := 6

	status := `REJECTED`

	if payment.Amount > highRiskAmount {
		payment.Status = &status
		as.saveDataToDatabase(payment)
		as.log.Info("Suspicious transaction: high amount, high risk - Email: %s - Amount: %.2f", zap.String("email", payment.Email), zap.Float64("amount", payment.Amount))
		return payment
	}

	if payment.Amount > suspiciousInstallmentAmount && payment.Installments > suspiciousInstallments {
		payment.Status = &status
		as.saveDataToDatabase(payment)
		as.log.Info("Suspicious transaction: high amount with elevated number of installments - Email: %s - Amount: %.2f - Installments: %d", zap.String("email", payment.Email), zap.Float64("amount", payment.Amount), zap.Int("installments", payment.Installments))
		return payment
	}

	if as.isSuspiciousName(payment.FirstName) || as.isSuspiciousName(payment.LastName) {
		payment.Status = &status
		as.saveDataToDatabase(payment)
		as.log.Info("Suspicious transaction: unusual name or surname - Email: %s - FirstName: %s - LastName: %s", zap.String("email", payment.Email), zap.String("first_name", payment.FirstName), zap.String("last_name", payment.LastName))
		return payment
	}

	status = "APPROVED"

	payment.Status = &status
	as.saveDataToDatabase(payment)

	return payment
}

func (as *AntifraudAnalisysService) isSuspiciousName(name string) bool {
	match, _ := regexp.MatchString(`^[A-Z]$`, name)
	return match
}

func (as *AntifraudAnalisysService) saveDataToDatabase(data topic_messages.Payment) (bool, error) {
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
