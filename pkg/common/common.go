package common

type Postgresql struct {
	HOST     string
	PORT     string
	USER     string
	PASSWORD string
	NAME     string
}

type Redis struct {
	ADDRESS      string
	ADDRESS_PORT string
}

type Prometheus struct {
	PUSHGATEWAY_URL string
}

type Mongo struct {
	URL     string
	DB_NAME string
}

type Kafka struct {
	URL          string
	OUTBOX_TOPIC string
}

type Grpc struct {
	ADDR      string
	ADDR_PORT string
}

type Steam struct {
	STEAM_KEY        string
	OWNGAME_ENDPOINT string
}
