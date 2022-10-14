package test

import (
	"bytes"
	"fmt"
	"github.com/cemayan/earthquake_collector_mini/config"
	"github.com/cemayan/earthquake_collector_mini/internal/service"
	kafka2 "github.com/cemayan/earthquake_collector_mini/pkg/kafka"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
	"golang.org/x/net/context"
	"os"
	"testing"
)

const (
	KAFKA_TOPIC = "last-earthquake"
)

type e2eTestSuite struct {
	suite.Suite
	configs      *config.AppConfig
	v            *viper.Viper
	kafkaCluster *KafkaCluster
	kafkaSvc     service.KafkaService
}

func (ts *e2eTestSuite) SetupSuite() {

	ts.v = viper.New()
	_configs := config.NewConfig(ts.v)

	env := os.Getenv("ENV")
	appConfig, err := _configs.GetConfig(env)
	ts.configs = appConfig
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(appConfig)

	kafkaHandler := kafka2.NewKafkaHandler(appConfig)
	ts.kafkaSvc = service.NewKafkaService(kafkaHandler)

	ts.kafkaSvc.KafkaProducer([]byte(`{"name":"2020.03.10 21:06:02","lokasyon":"MEDAR-AKHISAR (MANISA)                                            ","lat":"35.1457","lng":"27.9197","mag":"2.8","Depth":"3.3"}`))

}

func (ts *e2eTestSuite) TestConsumer() {
	eq := []byte(`{"name":"2020.03.10 21:06:02","lokasyon":"MEDAR-AKHISAR (MANISA)                                            ","lat":"35.1457","lng":"27.9197","mag":"2.8","Depth":"3.3"}`)
	fmt.Println(eq)

	// make a new reader that consumes from topic-A, partition 0, at offset 42
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{ts.configs.Kafka.URL},
		Topic:     ts.configs.Kafka.OUTBOX_TOPIC,
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})
	r.SetOffset(42)

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}

		compare := bytes.Compare(eq, m.Value)
		ts.Equal(compare, 0)
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}

}

func TestE2ETestSuite(t *testing.T) {
	suite.Run(t, &e2eTestSuite{})
}
