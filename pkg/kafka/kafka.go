package kafka

import (
	"github.com/cemayan/earthquake_collector_mini/config"
	"github.com/segmentio/kafka-go"
)

type KafkaHandler interface {
	GetKafkaWriter() *kafka.Writer
	GetKafkaReader() *kafka.Reader
}
type KafkaSvc struct {
	configs *config.AppConfig
}

func (k KafkaSvc) GetKafkaWriter() *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(k.configs.Kafka.URL),
		Topic:    k.configs.Kafka.OUTBOX_TOPIC,
		Balancer: &kafka.LeastBytes{},
	}
}

func (k KafkaSvc) GetKafkaReader() *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{k.configs.Kafka.URL},
		Topic:    k.configs.Kafka.OUTBOX_TOPIC,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
}

func NewKafkaHandler(configs *config.AppConfig) KafkaHandler {
	return &KafkaSvc{configs: configs}
}
