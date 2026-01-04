package repository

import (
	"context"
	"orders/model"

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

func (r *Repository) CreateOrder(ctx context.Context, order *model.Order) (*model.Order, error) {
	query := `INSERT INTO orders
		(user_id, amount, currency, payment_method, status)
		VALUES (@userID, @amount, @currency, @paymentMethod, @status)
		RETURNING *`

	args := pgx.NamedArgs{
		"userID":        order.UserID,
		"amount":        order.Amount,
		"currency":      order.Currency,
		"paymentMethod": order.PaymentMethod,
		"status":        order.Status,
	}

	err := r.pool.QueryRow(ctx, query, args).Scan(&order.ID, &order.UserID, &order.Status, &order.Currency, &order.Amount, &order.PaymentMethod, &order.CreatedAt, &order.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (r *Repository) CreateItems(ctx context.Context, items []model.OrderItem) error {
	rows := [][]any{}
	for _, item := range items {
		rows = append(rows, []any{item.OrderID, item.ProductID, item.ProductName, item.UnitPrice, item.Quantity, item.Amount})
	}

	tableName := pgx.Identifier{"order_items"}
	columns := []string{"order_id", "product_id", "product_name", "unit_price", "quantity", "amount"}

	if _, err := r.pool.CopyFrom(
		ctx,
		tableName,
		columns,
		pgx.CopyFromRows(rows),
	); err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateOrder(ctx context.Context, order *model.Order) (*model.Order, error) {
	query := `UPDATE orders set status = @status where id = @orderID RETURNING *`

	args := pgx.NamedArgs{
		"orderID": order.ID,
		"status":  order.Status,
	}

	err := r.pool.QueryRow(ctx, query, args).Scan(&order.ID, &order.UserID, &order.Status, &order.Currency, &order.Amount, &order.PaymentMethod, &order.CreatedAt, &order.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return order, nil
}
