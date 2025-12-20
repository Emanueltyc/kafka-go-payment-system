package handler

import (
	"net/http"
	"payment/dto"
	"payment/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type PaymentHandler struct {
	service *service.PaymentService
}

func NewPaymentHandler(service *service.PaymentService) *PaymentHandler {
	return &PaymentHandler{
		service: service,
	}
}

func (ph *PaymentHandler) Create(c *fiber.Ctx) error {
	var paymentRequest *dto.PaymentRequest

	if err := c.BodyParser(&paymentRequest); err != nil {
		return err
	}

	payment, err := ph.service.CreateOrder(c.Context(), paymentRequest)
	if err != nil {
		return err
	}

	paymentResponse := &dto.PaymentResponse{
		ID:            strconv.FormatInt(payment.ID, 10),
		UserID:        strconv.FormatInt(payment.UserID, 10),
		Status:        payment.Status,
		Currency:      payment.Currency,
		Amount:        payment.Amount,
		PaymentMethod: payment.PaymentMethod,
		CreatedAt:     payment.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	c.Status(http.StatusCreated).JSON(fiber.Map{
		"payment": paymentResponse,
	})

	return nil
}
