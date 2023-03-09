package RabbitMq

import (
	"BootCampT1/logger"
	"encoding/json"
	"github.com/streadway/amqp"
)

func Push(key string, data any) error {
	body, err := json.Marshal(data)
	if err != nil {
		logger.Error.Printf("Unable to Unmarshall obj: %+v", data)
		return err
	}
	err = channel.Publish(
		"",
		key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		},
	)
	if err != nil {
		return err
	}
	return nil
}
