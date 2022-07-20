package config

import "os"

// Config is representation of a OS Env values
type Config struct {
	SCHEDULE_UNIT     string
	SCHEDULE_INTERVAL string
	TOPIC_ID          string
	BOOTSTRAP_SERVER  string
	XML_ADDRESS       string
}

// GetConfig returns values based on given os.Getenv("ENV")
func GetConfig() Config {
	if os.Getenv("ENV") == "dev" {
		return Config{
			SCHEDULE_UNIT:     os.Getenv("SCHEDULE_UNIT_DEV"),
			SCHEDULE_INTERVAL: os.Getenv("SCHEDULE_INTERVAL_DEV"),
			TOPIC_ID:          os.Getenv("TOPIC_ID_DEV"),
			BOOTSTRAP_SERVER:  os.Getenv("BOOTSTRAP_SERVER_DEV"),
			XML_ADDRESS:       os.Getenv("XML_ADDRESS_DEV"),
		}
	} else {
		return Config{
			SCHEDULE_UNIT:     os.Getenv("SCHEDULE_UNIT_PROD"),
			SCHEDULE_INTERVAL: os.Getenv("SCHEDULE_INTERVAL_PROD"),
			TOPIC_ID:          os.Getenv("TOPIC_ID_PROD"),
			BOOTSTRAP_SERVER:  os.Getenv("BOOTSTRAP_SERVER_PROD"),
			XML_ADDRESS:       os.Getenv("XML_ADDRESS_PROD"),
		}
	}
}

// SetConfigForTesting sets the "ENV" value
func SetConfigForTesting() {
	os.Setenv("ENV", "test")
}
