package event

import "time"

type OrderCreated struct {
	EventID   string       `json:"event_id"`
	Payload   OrderPayload `json:"payload"`
	CreatedAt time.Time    `json:"created_at"`
}

type OrderPayload struct {
	OrderID  string      `json:"order_id"`
	UserID   string      `json:"user_id"`
	Username string      `json:"username"`
	Email    string      `json:"email"`
	Items    []OrderItem `json:"items"`
	Amount   string      `json:"amount"`
	Currency string      `json:"currency"`
}

type OrderItem struct {
	Name     string `json:"name"`
	Price    string `json:"price"`
	Quantity int    `json:"quantity"`
}
