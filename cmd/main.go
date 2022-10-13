package main

import (
	"github.com/cemayan/earthquake_collector_mini/config"
	service2 "github.com/cemayan/earthquake_collector_mini/internal/service"
	"github.com/cemayan/earthquake_collector_mini/pkg/bigcache"
	"github.com/cemayan/earthquake_collector_mini/pkg/kafka"
	"github.com/jasonlvhit/gocron"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"strconv"
)

var _log *logrus.Logger
var configs *config.AppConfig
var v *viper.Viper
var kafkaHandler kafka.KafkaHandler
var cacheHandler bigcache.BigcacheHandler

func init() {

	//logrus init
	_log = logrus.New()
	_log.Out = os.Stdout

	v = viper.New()
	_configs := config.NewConfig(v)

	env := os.Getenv("ENV")
	appConfig, err := _configs.GetConfig(env)
	configs = appConfig
	if err != nil {
		return
	}

	kafkaHandler = kafka.NewKafkaHandler(configs)
	cacheHandler = bigcache.NewBigcacheHandler()

}

func main() {

	cacheManager := cacheHandler.New()
	scheduler := gocron.NewScheduler()

	interval, _ := strconv.ParseUint(configs.SCHEDULE_INTERVAL, 0, 64)

	var xmlSvc service2.XMLService
	xmlSvc = service2.NewXMLService()

	var kafkaSvc service2.KafkaService
	kafkaSvc = service2.NewKafkaService(kafkaHandler)

	var schedulerSvc service2.ScheduleService
	schedulerSvc = service2.NewSchedulerService(cacheManager, xmlSvc, kafkaSvc, configs)

	if configs.SCHEDULE_UNIT == "SECOND" {
		scheduler.Every(interval).Second().Do(schedulerSvc.ScheduleJob)
	} else if configs.SCHEDULE_UNIT == "MINUTE" {
		scheduler.Every(interval).Minute().Do(schedulerSvc.ScheduleJob)
	}

	<-scheduler.Start()
}
