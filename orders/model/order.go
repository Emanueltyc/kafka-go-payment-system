package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Order struct {
	ID            int64           `db:"id" json:"id,omitempty"`
	UserID        int64           `db:"user_id" json:"user_id"`
	Status        string          `db:"status" json:"status,omitempty"`
	Currency      string          `db:"currency" json:"currency"`
	Amount        decimal.Decimal `db:"amount" json:"amount"`
	PaymentMethod string          `db:"payment_method" json:"payment_method"`
	CreatedAt     time.Time       `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time       `db:"updated_at" json:"updated_at"`
}
