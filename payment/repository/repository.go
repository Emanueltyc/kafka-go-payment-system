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

func (r *Repository) UpdatePayment(ctx context.Context, payment *model.Payment) (*model.Payment, error) {
	query := `UPDATE payments SET status = @status WHERE ID = @PaymentID RETURNING *`

	args := pgx.NamedArgs{
		"PaymentID": payment.ID,
		"status":    payment.Status,
	}

	err := r.pool.QueryRow(ctx, query, args).Scan(&payment.ID, &payment.OrderID, &payment.UserID, &payment.Status, &payment.Currency, &payment.Amount, &payment.PaymentMethod, &payment.CreatedAt, &payment.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (r *Repository) CreateProcessedEvent(ctx context.Context, event *model.ProcessedEvent) error {
	query := `INSERT INTO processed_events (id) VALUES (@id)`

	if _, err := r.pool.Exec(ctx, query, pgx.NamedArgs{
		"id": event.ID,
	}); err != nil {
		return err
	}

	return nil
}

func (r *Repository) FindProcessedEventByID(ctx context.Context, id string) (*model.ProcessedEvent, error) {
	query := `SELECT * FROM processed_events WHERE id = @id`

	processedEvent := new(model.ProcessedEvent)

	if err := r.pool.QueryRow(ctx, query, pgx.NamedArgs{
		"id": id,
	}).Scan(&processedEvent.ID, &processedEvent.CreatedAt); err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		
		return nil, err
	}

	return processedEvent, nil
}