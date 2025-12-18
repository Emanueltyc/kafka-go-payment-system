package handler

import (
	"net/http"
	"orders/dto"
	"orders/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
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
