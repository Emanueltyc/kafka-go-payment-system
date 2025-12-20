package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Payment struct {
	ID            int64           `db:"id"`
	OrderID       int64           `db:"order_id"`
	UserID        int64           `db:"user_id"`
	Amount        decimal.Decimal `db:"amount"`
	Currency      string          `db:"currency"`
	PaymentMethod string          `db:"payment_method"`
	Status        string          `db:"status"`
	CreatedAt     time.Time       `db:"created_at"`
	UpdatedAt     time.Time       `db:"updated_at"`
}
