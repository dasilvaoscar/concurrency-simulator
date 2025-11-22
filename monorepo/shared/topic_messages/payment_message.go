package topic_messages

type Payment struct {
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	Amount       float64 `json:"amount"` // Só aceita valor com até 8 dígitos
	Installments int     `json:"installments"`
	Email        string  `json:"email"`
	Status       *string
}
