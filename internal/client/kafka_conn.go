package client

import (
	"context"
	"github.com/cemayan/earthquake_collector_mini/config"
	"github.com/segmentio/kafka-go"
	"log"
)

func NewKafkaClient(configs *config.AppConfig) *kafka.Conn {

	conn, err := kafka.DialLeader(context.Background(), "tcp",
		configs.BOOTSTRAP_SERVER, configs.TOPIC_ID, 0)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	return conn
}
