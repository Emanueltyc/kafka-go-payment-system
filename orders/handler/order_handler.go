package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"orders/dto"
	"orders/event"
	"orders/model"
	"orders/service"
	"orders/status"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/segmentio/kafka-go"
)

type OrderHandler struct {
	service *service.OrderService
}

func NewOrderHandler(service *service.OrderService) *OrderHandler {
	return &OrderHandler{
		service: service,
	}
}

func (oc *OrderHandler) Create(c *fiber.Ctx) error {
	var orderRequest *dto.OrderRequest

	if err := c.BodyParser(&orderRequest); err != nil {
		return err
	}

	order, err := oc.service.CreateOrder(c.Context(), orderRequest)
	if err != nil {
		return err
	}

	orderResponse := &dto.OrderResponse{
		ID:            strconv.FormatInt(order.ID, 10),
		UserID:        strconv.FormatInt(order.UserID, 10),
		Status:        order.Status,
		Currency:      order.Currency,
		Amount:        order.Amount,
		PaymentMethod: order.PaymentMethod,
		CreatedAt:     order.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	c.Status(http.StatusCreated).JSON(fiber.Map{
		"order": orderResponse,
	})

	return nil
}

func (oc *OrderHandler) HandleMessage(m kafka.Message) error {
	var paymentEvent event.PaymentEvent

	if err := json.Unmarshal(m.Value, &paymentEvent); err != nil {
		return err
	}

	orderID, _ := strconv.ParseInt(paymentEvent.OrderID, 10, 64)

	var newStatus string
	switch m.Topic {
	case "payment.approved":
		newStatus = status.APPROVED
	case "payment.rejected":
		newStatus = status.REJECTED
	}

	order := &model.Order{
		ID:     orderID,
		Status: newStatus,
	}

	if _, err := oc.service.UpdateOrder(context.Background(), order); err != nil {
		return err
	}

	return nil
}
