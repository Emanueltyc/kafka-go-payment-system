package ports

import "context"

type ChargeRequest struct {
	PaymentID string
	PaymentToken string
	Amount       string
	Currency     string
	OrderID      string
	UserID       string
}

type ChargeResult struct {
	PaymentID string
	Status    string
	Reason    string
}

type PaymentGateway interface {
	Charge(ctx context.Context, req *ChargeRequest) (*ChargeResult, error)
}
