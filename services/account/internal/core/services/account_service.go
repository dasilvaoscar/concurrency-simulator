package services

import (
	"concurrency-simulator/services/shared/topic_messages"
	"database/sql"

	"go.uber.org/zap"
)

type AccountService struct {
	log    *zap.Logger
	driver *sql.DB
}

func (ac *AccountService) Execute(payment topic_messages.Payment) bool {
	tx, err := ac.driver.Begin()
	if err != nil {
		ac.log.Error("failed to ping database", zap.Error(err))
		return false
	}
	defer tx.Rollback()

	is_account_created := ac.accountExists(tx, payment.Email)

	if is_account_created {
		ac.log.Info("Account already exists", zap.String("email", payment.Email))
		return is_account_created
	}

	ac.log.Info("Account not found, creating account", zap.String("email", payment.Email))
	is_account_created = ac.createAccount(tx, payment)

	if err := tx.Commit(); err != nil {
		ac.log.Error("failed to commit transaction", zap.Error(err))
		return false
	}

	return is_account_created
}

func (ac *AccountService) accountExists(tx *sql.Tx, email string) bool {
	query := `
		SELECT EXISTS (
			SELECT 1 FROM account
			WHERE email = $1
		)
	`

	var exists bool
	err := tx.QueryRow(query, email).Scan(&exists)
	if err != nil {
		ac.log.Error("failed to check account exists", zap.Error(err), zap.String("email", email))
		return false
	}

	ac.log.Info("QUERY COMPLETED", zap.String("email", email), zap.Bool("exists", exists))
	return exists
}

func (ac *AccountService) createAccount(tx *sql.Tx, data topic_messages.Payment) bool {
	query := `
		INSERT INTO account (first_name, last_name, email, created_at) 
		VALUES ($1, $2, $3, NOW())
	`

	_, err := tx.Exec(query, data.FirstName, data.LastName, data.Email)

	if err != nil {
		ac.log.Error("failed to create account", zap.Error(err), zap.String("email", data.Email))
		return false
	}

	return true
}
