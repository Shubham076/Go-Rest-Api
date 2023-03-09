package RabbitMq

import (
	"BootCampT1/config"
	"BootCampT1/logger"
	"github.com/streadway/amqp"
)

var conn *amqp.Connection
var channel *amqp.Channel

func createUrl() string {
	conf := config.GetConfig()
	url := "amqps://" + conf.Mq.User + ":" + conf.Mq.Pass + "@" + conf.Mq.Host
	return url
}

func Init() {
	var err error
	conn, err = amqp.Dial(createUrl())
	channel, err = conn.Channel()
	if err != nil {
		logger.Error.Printf("Can't connect to Rabbit Mq err: %s", err.Error())
		return
	}
	logger.Info.Println("Connected to Rabbit Mq instance")
	initQueues()
	initConsumers()
}
