package service

import (
	"context"
	"log"
	"payment/constants"
	"payment/dto"
	"payment/event"
	"payment/model"
	"payment/repository"
	"payment/streaming"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type PaymentService struct {
	repo     *repository.Repository
	producer *streaming.Producer
}

func NewPaymentService(repo *repository.Repository, producer *streaming.Producer) *PaymentService {
	return &PaymentService{
		repo:     repo,
		producer: producer,
	}
}

func (s *PaymentService) CreateOrder(ctx context.Context, paymentRequest *dto.PaymentRequest) (*model.Payment, error) {
	orderID, _ := strconv.ParseInt(paymentRequest.OrderID, 10, 64)
	userID, _ := strconv.ParseInt(paymentRequest.UserID, 10, 64)
	amount, _ := decimal.NewFromString(paymentRequest.Amount)

	/*
		to do: payment gateway mock call
	*/

	payment := &model.Payment{
		OrderID:       orderID,
		UserID:        userID,
		Currency:      paymentRequest.Currency,
		Amount:        amount,
		PaymentMethod: paymentRequest.PaymentMethod,
		Status:        constants.Status["pending"],
	}

	payment, err := s.repo.CreatePayment(ctx, payment)
	if err != nil {
		return nil, err
	}

	event := &event.PaymentEvent{
		EventID:   uuid.NewString(),
		PaymentID: strconv.FormatInt(payment.ID, 10),
		OrderID:   strconv.FormatInt(payment.OrderID, 10),
		CreatedAt: time.Now(),
	}

	go func() {
		if err := s.producer.Publish(ctx, constants.Topics["pending"], event); err != nil {
			log.Println(err)
		}
	}()

	return payment, nil
}
