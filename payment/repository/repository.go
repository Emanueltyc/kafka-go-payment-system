package repository

import (
	"context"
	"payment/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		pool: pool,
	}
}

func (r *Repository) CreatePayment(ctx context.Context, payment *model.Payment) (*model.Payment, error) {
	query := `INSERT INTO payments
		(order_id, user_id, amount, currency, payment_method, status)
		VALUES (@orderID, @userID, @amount, @currency, @paymentMethod, @status)
		RETURNING *`

	args := pgx.NamedArgs{
		"orderID":       payment.OrderID,
		"userID":        payment.UserID,
		"amount":        payment.Amount,
		"currency":      payment.Currency,
		"paymentMethod": payment.PaymentMethod,
		"status":        payment.Status,
	}

	err := r.pool.QueryRow(ctx, query, args).Scan(&payment.ID, &payment.OrderID, &payment.UserID, &payment.Status, &payment.Currency, &payment.Amount, &payment.PaymentMethod, &payment.CreatedAt, &payment.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return payment, nil
}
