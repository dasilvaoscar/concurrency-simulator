package services

import (
	"concurrency-simulator/monorepo/shared/topic_messages"
	"database/sql"

	"go.uber.org/zap"
)

type AccountService struct {
	log    *zap.Logger
	driver *sql.DB
}

func (ac *AccountService) Execute(payment topic_messages.Payment) bool {
	is_account_created := ac.accountExists(payment.Email)

	if is_account_created {
		return is_account_created
	}

	is_account_created = ac.createAccount(payment)
	return is_account_created
}

func (ac *AccountService) accountExists(email string) bool {
	query := `
		SELECT EXISTS (
			SELECT 1 FROM account
			WHERE email = $1
		)
	`

	var exists bool
	err := ac.driver.QueryRow(query, email).Scan(&exists)
	if err != nil {
		ac.log.Error("failed to check account exists", zap.Error(err), zap.String("email", email))
		return false
	}

	return exists
}

func (ac *AccountService) createAccount(data topic_messages.Payment) bool {
	query := `
		INSERT INTO account (first_name, last_name, email, created_at) 
		VALUES ($1, $2, $3, NOW())
	`

	_, err := ac.driver.Exec(query, data.FirstName, data.LastName, data.Email)

	if err != nil {
		ac.log.Fatal(err.Error())
		panic("save-data-error")
	}

	return true
}
