package streaming

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader  *kafka.Reader
	handler func(m kafka.Message) error
}

func NewConsumer(addr string, handler func(m kafka.Message) error) *Consumer {
	return &Consumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:     []string{addr},
			GroupTopics: []string{"order.created", "payment.approved", "payment.rejected"},
			GroupID:     "notification",
		}),
		handler: handler,
	}
}

func (c *Consumer) Read(ctx context.Context) {
	for {
		message, err := c.reader.ReadMessage(ctx)
		if err != nil {
			log.Println("erro no consumer: ", err)
			break
		}

		go c.handler(message)
	}


	c.reader.Close()
}
