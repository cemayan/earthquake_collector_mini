package main

import (
	"github.com/cemayan/earthquake_collector_mini/src/client"
	"github.com/cemayan/earthquake_collector_mini/src/config"
	"github.com/cemayan/earthquake_collector_mini/src/service"
	"github.com/jasonlvhit/gocron"
	"strconv"
)

func main() {

	configs := config.GetConfig()
	cacheManager := client.InitCache()
	scheduler := gocron.NewScheduler()
	kafkaConn := client.NewKafkaClient()

	interval, _ := strconv.ParseUint(configs.SCHEDULE_INTERVAL, 0, 64)

	var xmlSvc service.XMLService
	xmlSvc = service.NewXMLService()

	var kafkaSvc service.KafkaService
	kafkaSvc = service.NewKafkaService(kafkaConn)

	var schedulerSvc service.ScheduleService
	schedulerSvc = service.NewSchedulerService(cacheManager, xmlSvc, kafkaSvc)

	if configs.SCHEDULE_UNIT == "SECOND" {
		scheduler.Every(interval).Second().Do(schedulerSvc.ScheduleJob)
	} else if configs.SCHEDULE_UNIT == "MINUTE" {
		scheduler.Every(interval).Minute().Do(schedulerSvc.ScheduleJob)
	}

	<-scheduler.Start()
}
