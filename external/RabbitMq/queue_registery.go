package RabbitMq

type QueueData struct {
	QueueName            string
	Exchange             string
	RoutingKey           string
	DeadLetterEnabled    bool
	DeadLetterExchange   string
	DeadLetterRoutingKey string
	Prefetch             int
}

var queues = map[string]*QueueData{
	"Email": &QueueData{
		QueueName:          "Email",
		Exchange:           "",
		RoutingKey:         "Email",
		DeadLetterEnabled:  true,
		DeadLetterExchange: "secretDetector.sideline",
		Prefetch:           1,
	},
	"DeleteSession": &QueueData{
		QueueName:          "DeleteSession",
		Exchange:           "",
		RoutingKey:         "DeleteSession",
		DeadLetterEnabled:  true,
		DeadLetterExchange: "secretDetector.sideline",
		Prefetch:           1,
	},
}
