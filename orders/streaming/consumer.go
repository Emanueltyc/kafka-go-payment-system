package streaming

import (
	"context"

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
			GroupTopics: []string{"payment.approved", "payment.rejected"},
			GroupID:     "payment",
		}),
		handler: handler,
	}
}

func (c *Consumer) Read(ctx context.Context) {
	for {
		message, err := c.reader.ReadMessage(ctx)
		if err != nil {
			break
		}

		go c.handler(message)
	}

	c.reader.Close()
}
