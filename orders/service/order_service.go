package service

import (
	"context"
	"log"
	"orders/dto"
	"orders/event"
	"orders/model"
	"orders/repository"
	"orders/status"
	"orders/streaming"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type OrderService struct {
	repo     *repository.Repository
	producer *streaming.Producer
}

func NewOrderService(repo *repository.Repository, producer *streaming.Producer) *OrderService {
	return &OrderService{
		repo:     repo,
		producer: producer,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, orderRequest *dto.OrderRequest) (*model.Order, error) {
	userID, _ := strconv.ParseInt(orderRequest.UserID, 10, 64)
	amount, _ := decimal.NewFromString(orderRequest.Amount)

	order := &model.Order{
		UserID:        userID,
		Currency:      orderRequest.Currency,
		Amount:        amount,
		PaymentMethod: orderRequest.PaymentMethod,
		Status:        status.CREATED,
	}

	order, err := s.repo.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	var items []model.OrderItem

	for _, item := range orderRequest.Items {
		productID, _ := strconv.ParseInt(item.ProductID, 10, 64)
		unitPrice, _ := decimal.NewFromString(item.UnitPrice)
		amount, _ := decimal.NewFromString(item.Amount)

		items = append(items, model.OrderItem{
			OrderID:     order.ID,
			ProductID:   productID,
			ProductName: item.Name,
			UnitPrice:   unitPrice,
			Quantity:    int32(item.Quantity),
			Amount:      amount,
		})
	}

	if err = s.repo.CreateItems(ctx, items); err != nil {
		return nil, err
	}

	newEvent := &event.OrderCreated{
		EventID: uuid.NewString(),
		Payload: event.OrderPayload{
			OrderID:  strconv.FormatInt(order.ID, 10),
			Username: orderRequest.Username,
			Email:    orderRequest.Email,
			UserID:   strconv.FormatInt(order.UserID, 10),
			Amount:   order.Amount.String(),
			Currency: order.Currency,
			Items:    []event.OrderItem{},
		},
		CreatedAt: time.Now(),
	}

	for _, i := range items {
		newEvent.Payload.Items = append(newEvent.Payload.Items, event.OrderItem{
			Name:     i.ProductName,
			Price:    i.UnitPrice.String(),
			Quantity: int(i.Quantity),
		})
	}

	go func() {
		if err := s.producer.Publish(ctx, newEvent); err != nil {
			log.Println(err)
		}
	}()

	return order, nil
}

func (s *OrderService) CreateItems(ctx context.Context, items []model.OrderItem) error {
	if err := s.repo.CreateItems(ctx, items); err != nil {
		return err
	}

	return nil
}

func (s *OrderService) UpdateOrder(ctx context.Context, order *model.Order) (*model.Order, error) {
	order, err := s.repo.UpdateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	return order, nil
}
