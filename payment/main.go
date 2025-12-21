package main

import (
	"log"
	"os"
	"payment/database"
	"payment/gateway"
	"payment/handler"
	"payment/middleware"
	"payment/repository"
	"payment/router"
	"payment/service"
	"payment/streaming"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	db := database.Connect()
	defer db.Close()

	streaming.CreateTopics(os.Getenv("KAFKA_ADDR"))
	producer := streaming.NewProducer(os.Getenv("KAFKA_ADDR"))

	repository := repository.NewRepository(db)
	gateway := gateway.NewMockGateway()
	paymentService := service.NewPaymentService(repository, producer, &gateway)
	paymentHandler := handler.NewPaymentHandler(paymentService)

	api := app.Group("/api/v1", middleware.JWTValidator)
	api.Route("/", func(r fiber.Router) {
		router.PaymentRouter(r, paymentHandler)
	})

	webhook := app.Group("/webhook/v1", middleware.SignatureValidator)
	webhook.Route("/", func(r fiber.Router) {
		router.WebhookRouter(r, paymentHandler)
	})

	log.Fatal(app.Listen(":8080"))
}
