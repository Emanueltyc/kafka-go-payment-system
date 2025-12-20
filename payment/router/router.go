package router

import (
	"payment/handler"

	"github.com/gofiber/fiber/v2"
)

func PaymentRouter(r fiber.Router, ph *handler.PaymentHandler) {
	r.Post("/payment", ph.Create)
}
