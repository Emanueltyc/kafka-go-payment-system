package main

import (
	"context"
	"notification/email"
	"notification/handler"
	"notification/streaming"
	"os"
	"strconv"
)

func main() {
	smtpPort, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))

	sender := &email.SmtpSender{
		Host: os.Getenv("SMTP_HOST"),
		Port: smtpPort,
	}
	
	renderer, _ := email.NewRenderer()

	notificationHandler := handler.NewNotificationHandler(sender, renderer)

	consumer := streaming.NewConsumer(os.Getenv("KAFKA_ADDR"), notificationHandler.HandleMessage)
	consumer.Read(context.Background())
}
