package event

import "time"

type PaymentEvent struct {
	EventID   string    `json:"event_id"`
	PaymentID string    `json:"payment_id"`
	OrderID   string    `json:"order_id"`
	CreatedAt time.Time `json:"created_at"`
}
