package model

import "github.com/shopspring/decimal"

type OrderItem struct {
	ID          int64           `db:"id" json:"id"`
	OrderID     int64           `db:"order_id" json:"order_id"`
	ProductID   int64           `db:"product_id" json:"product_id"`
	ProductName string          `db:"product_name" json:"product_name"`
	UnitPrice   decimal.Decimal `db:"unit_price" json:"unit_price"`
	Quantity    int32           `db:"quantity" json:"quantity"`
	Amount      decimal.Decimal `db:"amount" json:"amount"`
	CreatedAt   string          `db:"created_at" json:"created_at,omitempty"`
}
