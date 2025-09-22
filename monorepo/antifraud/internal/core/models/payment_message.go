package models

type PaymentMessage struct {
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	Amount       float64 `json:"amount"`
	Installments int     `json:"installments"`
	Email        string  `json:"email"`
}
