package dto

import "github.com/shopspring/decimal"

type PaymentRequest struct {
	OrderID       string `json:"order_id" validate:"required"`
	UserID        string `json:"user_id" validate:"required"`
	Amount        string `json:"amount" validate:"required"`
	Currency      string `json:"currency" validate:"required"`
	PaymentMethod string `json:"payment_method" validate:"required"`
	PaymentToken  string `json:"payment_token" validate:"required"`
}

type PaymentResponse struct {
	ID            string          `json:"id"`
	UserID        string          `json:"user_id"`
	Status        string          `json:"status"`
	Currency      string          `json:"currency"`
	Amount        decimal.Decimal `json:"amount"`
	PaymentMethod string          `json:"payment_method"`
	CreatedAt     string          `json:"created_at"`
}
