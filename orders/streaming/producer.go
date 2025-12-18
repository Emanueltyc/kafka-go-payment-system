package streaming

import (
	"context"
	"encoding/json"
	"log"
	"orders/event"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(addr string, topic string) *Producer {
	conn, err := kafka.Dial("tcp", addr)
	if err != nil {
		log.Println(err)
	}
	
	if err = conn.CreateTopics(kafka.TopicConfig{
		Topic: topic,
		NumPartitions: 1,
		ReplicationFactor: 1,
	}); err != nil {
		log.Println(err)
	}

	defer conn.Close()

	return &Producer{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(addr),
			Topic:    topic,
			Balancer: &kafka.LeastBytes{},
			AllowAutoTopicCreation: true,
		},
	}
}

func (p *Producer) Publish(ctx context.Context, event *event.OrderCreated) error {
	bytes, err := json.Marshal(event)
	if err != nil {
		return err
	}
	
	return p.writer.WriteMessages(ctx, kafka.Message{
		Value: bytes,
	})
}
