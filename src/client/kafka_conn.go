package client

import (
	"context"
	"github.com/cemayan/earthquake_collector_mini/src/config"
	"github.com/segmentio/kafka-go"
	"log"
)

func NewKafkaClient() *kafka.Conn {

	configs := config.GetConfig()

	conn, err := kafka.DialLeader(context.Background(), "tcp",
		configs.BOOTSTRAP_SERVER, configs.TOPIC_ID, 0)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	return conn
}
