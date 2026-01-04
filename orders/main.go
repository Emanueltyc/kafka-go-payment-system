package main

import (
	"context"
	"log"
	"orders/database"
	"orders/handler"
	"orders/middleware"
	"orders/repository"
	"orders/router"
	"orders/service"
	"orders/streaming"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	db := database.Connect()
	defer db.Close()

	producer := streaming.NewProducer(os.Getenv("KAFKA_ADDR"), "order.created")

	orderRepo := repository.NewRepository(db)
	orderService := service.NewOrderService(orderRepo, producer)
	orderHandler := handler.NewOrderHandler(orderService)

	consumer := streaming.NewConsumer(os.Getenv("KAFKA_ADDR"), orderHandler.HandleMessage)
	go consumer.Read(context.Background())

	api := app.Group("/api/v1", middleware.Validator)

	api.Route("/", func(r fiber.Router) {
		router.OrderRouter(r, orderHandler)
	})

	log.Fatal(app.Listen(":8080"))
}
