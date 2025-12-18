package router

import (
	"orders/handler"

	"github.com/gofiber/fiber/v2"
)

func OrderRouter(r fiber.Router, oc *handler.OrderHandler) {
	r.Post("/order", oc.Create)
}
