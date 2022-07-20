package service

type KafkaService interface {
	KafkaProducer(lastEq []byte)
}
