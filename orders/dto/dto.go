package dto

import "github.com/shopspring/decimal"

type OrderRequest struct {
	UserID        string             `json:"user_id" validate:"required"`
	Username      string             `json:"username" validate:"required"`
	Email         string             `json:"email" validate:"required"`
	Items         []OrderItemRequest `json:"items" validate:"required"`
	Currency      string             `json:"currency" validate:"required"`
	Amount        string             `json:"amount" validate:"required"`
	PaymentMethod string             `json:"payment_method" validate:"required"`
}

type OrderItemRequest struct {
	ProductID string `json:"product_id"`
	Name      string `json:"name"`
	UnitPrice string `json:"unit_price"`
	Quantity  int    `json:"quantity"`
	Amount    string `json:"Amount"`
}

type OrderResponse struct {
	ID            string          `json:"id"`
	UserID        string          `json:"user_id"`
	Status        string          `json:"status,omitempty"`
	Currency      string          `json:"currency"`
	Amount        decimal.Decimal `json:"amount"`
	PaymentMethod string          `json:"payment_method"`
	CreatedAt     string          `json:"created_at,omitempty"`
}
