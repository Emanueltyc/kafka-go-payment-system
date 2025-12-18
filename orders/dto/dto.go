package dto

import "github.com/shopspring/decimal"

type OrderRequest struct {
	UserID        string             `json:"user_id"`
	Items         []OrderItemRequest `json:"items"`
	Currency      string             `json:"currency"`
	Amount        string             `json:"amount"`
	PaymentMethod string             `json:"payment_method"`
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
