package handler

import (
	"encoding/json"
	"notification/dto"
	"notification/email"
	"notification/event"
	"strings"

	"github.com/segmentio/kafka-go"
)

type NotificationHandler struct {
	sender   *email.SmtpSender
	Renderer *email.Renderer
}

func NewNotificationHandler(sender *email.SmtpSender, renderer *email.Renderer) *NotificationHandler {
	return &NotificationHandler{
		sender:   sender,
		Renderer: renderer,
	}
}

var currencies map[string]string = map[string]string{
	"USD": "$",
	"EUR": "€",
	"BRL": "R$",
}

func (nh *NotificationHandler) HandleMessage(m kafka.Message) error {
	html := ""
	email := ""
	subject := ""

	switch strings.Split(m.Topic, ".")[0] {
	case "order":
		var orderEvent event.OrderCreated

		if err := json.Unmarshal(m.Value, &orderEvent); err != nil {
			return err
		}

		data := dto.OrderCreatedTemplate{
			Title:    "Order created",
			Username: orderEvent.Payload.Username,
			OrderID:  orderEvent.Payload.OrderID,
			Amount:   orderEvent.Payload.Amount,
			Currency: currencies[orderEvent.Payload.Currency],
			Items:    []dto.OrderItemTemplate{},
		}

		for _, i := range orderEvent.Payload.Items {
			data.Items = append(data.Items, dto.OrderItemTemplate{Name: i.Name, Quantity: i.Quantity, Price: i.Price})
		}

		renderedString, err := nh.Renderer.Render("order_created", data)
		if err != nil {
			return err
		}

		html = renderedString
		email = orderEvent.Payload.Email
		subject = "Order created successfully"
	case "payment":
		// to implement
	}

	if err := nh.sender.Send(email, subject, html); err != nil {
		return err
	}

	return nil
}
