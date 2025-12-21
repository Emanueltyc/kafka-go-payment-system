package service

import (
	"context"
	"log"
	"payment/constants"
	"payment/dto"
	"payment/event"
	"payment/model"
	"payment/ports"
	"payment/repository"
	"payment/streaming"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type PaymentService struct {
	repo     *repository.Repository
	producer *streaming.Producer
	gateway  ports.PaymentGateway
}

func NewPaymentService(repo *repository.Repository, producer *streaming.Producer, gateway ports.PaymentGateway) *PaymentService {
	return &PaymentService{
		repo:     repo,
		producer: producer,
		gateway:  gateway,
	}
}

func (s *PaymentService) CreatePayment(ctx context.Context, paymentRequest *dto.PaymentRequest) (*model.Payment, error) {
	orderID, _ := strconv.ParseInt(paymentRequest.OrderID, 10, 64)
	userID, _ := strconv.ParseInt(paymentRequest.UserID, 10, 64)
	amount, _ := decimal.NewFromString(paymentRequest.Amount)

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

	chargeRequest := &ports.ChargeRequest{
		PaymentID:    strconv.FormatInt(payment.ID, 10),
		OrderID:      paymentRequest.OrderID,
		UserID:       paymentRequest.UserID,
		PaymentToken: paymentRequest.PaymentToken,
		Amount:       paymentRequest.Amount,
		Currency:     paymentRequest.Currency,
	}

	go s.gateway.Charge(ctx, chargeRequest)

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

func (s *PaymentService) UpdatePayment(ctx context.Context, gatewayRequest *dto.GatewayRequest) error {
	paymentID, _ := strconv.ParseInt(gatewayRequest.Data.PaymentID, 10, 64)

	status := constants.Status[strings.ToLower(gatewayRequest.Data.Status)]
	topic := constants.Topics[strings.ToLower(gatewayRequest.Data.Status)]

	payment := &model.Payment{
		ID:     paymentID,
		Status: status,
	}

	payment, err := s.repo.UpdatePayment(ctx, payment)
	if err != nil {
		return err
	}

	event := &event.PaymentEvent{
		EventID:   uuid.NewString(),
		PaymentID: strconv.FormatInt(payment.ID, 10),
		OrderID:   strconv.FormatInt(payment.OrderID, 10),
		CreatedAt: time.Now(),
	}

	go func() {
		if err := s.producer.Publish(ctx, topic, event); err != nil {
			log.Println(err)
		}
	}()

	return nil
}
