package streaming

import (
	"log"
	"payment/constants"

	"github.com/segmentio/kafka-go"
)

func CreateTopics(addr string) {
	conn, err := kafka.Dial("tcp", addr)
	if err != nil {
		log.Println(err)
	}
	
	topicConfigs := []kafka.TopicConfig{}

	for _, topic := range constants.Topics {
		topicConfigs = append(topicConfigs, kafka.TopicConfig{
			Topic:             topic,
			NumPartitions:     1,
			ReplicationFactor: 1,
		})
	}

	if err = conn.CreateTopics(topicConfigs...); err != nil {
		log.Println(err)
	}

	defer conn.Close()
}