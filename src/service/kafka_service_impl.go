package service

import (
	"github.com/segmentio/kafka-go"
	"log"
)

type KafkaSvc struct {
	kafkaConn *kafka.Conn
}

func (k KafkaSvc) KafkaProducer(lastEq []byte) {
	_, err := k.kafkaConn.Write(lastEq)

	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := k.kafkaConn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}

func NewKafkaService(kafkaConn *kafka.Conn) KafkaService {
	return &KafkaSvc{kafkaConn: kafkaConn}
}
