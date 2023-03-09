package RabbitMq

import (
	"BootCampT1/external/RabbitMq/handlers"
	"BootCampT1/logger"
	"github.com/streadway/amqp"
)

func createConsumers(no int, key string, handler func(msg amqp.Delivery)) {

	for i := 0; i < no; i++ {
		go func() {
			ch, err := conn.Channel()
			if err != nil {
				logger.Error.Println("Failed to create channel for consumer")
			}
			msgs, err := ch.Consume(
				key,
				"",
				true,
				false,
				false,
				false,
				nil,
			)

			if err != nil {
				logger.Error.Printf("Failed to consume messages from the queue: %s", key)
				return
			}

			for msg := range msgs {
				handler(msg)
			}
		}()
	}
}

func initConsumers() {
	createConsumers(2, "Email", handlers.SendMailHandler)
	createConsumers(2, "DeleteSession", handlers.DeleteSessionHandler)
}
