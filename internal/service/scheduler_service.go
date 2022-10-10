package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/allegro/bigcache"
	"github.com/cemayan/earthquake_collector_mini/config"
)

// ScheduleJob comment
type ScheduleService interface {
	ScheduleJob()
}

type SchedulerSvc struct {
	cacheManager *bigcache.BigCache
	xmlService   XMLService
	kafkaService KafkaService
	configs      *config.AppConfig
}

func (s SchedulerSvc) ScheduleJob() {

	data := s.xmlService.Fetch(s.configs.XML_ADDRESS)
	fmt.Println("Operation completed!")
	parsedData := s.xmlService.Parse(data)

	eq, _ := json.Marshal(parsedData[len(parsedData)-1])
	lastEarthQuake, _ := s.cacheManager.Get("last-earthquake")

	res := bytes.Compare(eq, lastEarthQuake)

	if res != 0 {
		s.cacheManager.Set("last-earthquake", eq)
		s.kafkaService.KafkaProducer(eq)
	} else {
		fmt.Println("there is no current earthquake right now")
	}
}

func NewSchedulerService(cacheManager *bigcache.BigCache, xmlService XMLService, kafkaService KafkaService, configs *config.AppConfig) ScheduleService {
	return &SchedulerSvc{cacheManager: cacheManager, xmlService: xmlService, kafkaService: kafkaService, configs: configs}
}
