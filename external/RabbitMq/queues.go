package RabbitMq

import "BootCampT1/logger"

func initQueues() {
	for k := range queues {
		_, err := channel.QueueDeclare(
			queues[k].QueueName, // name
			false,               // durable
			false,               // auto delete
			false,               // exclusive
			false,               // no wait
			nil,                 // args
		)
		if err != nil {
			logger.Error.Printf("Failed to create queue: %s")
		}
	}
}
