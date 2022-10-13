package service

import (
	"github.com/cemayan/earthquake_collector_mini/pkg/kafka"
	"github.com/google/uuid"
	kafka_ "github.com/segmentio/kafka-go"
	"golang.org/x/net/context"
	"log"
)

type KafkaService interface {
	KafkaProducer(lastEq []byte)
}

type KafkaSvc struct {
	kafkaHandler kafka.KafkaHandler
}

func (k KafkaSvc) KafkaProducer(lastEq []byte) {
	kafkaWriter := k.kafkaHandler.GetKafkaWriter()
	defer kafkaWriter.Close()

	msg := kafka_.Message{
		Key:   []byte(uuid.New().String()),
		Value: lastEq,
	}
	err := kafkaWriter.WriteMessages(context.Background(), msg)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}
}

func NewKafkaService(kafkaHandler kafka.KafkaHandler) KafkaService {
	return &KafkaSvc{kafkaHandler: kafkaHandler}
}
