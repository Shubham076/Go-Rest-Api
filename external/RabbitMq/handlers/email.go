package handlers

import (
	"BootCampT1/external/ses"
	"BootCampT1/logger"
	"encoding/json"
	"github.com/streadway/amqp"
)

func SendMailHandler(msg amqp.Delivery) {
	var req SendMailRequest
	err := json.Unmarshal(msg.Body, &req)

	if err != nil {
		logger.Error.Printf("Unable to unmarhshall msg due to err: %s", err.Error())
		return
	}

	ses.SendMail(req.EmailConf)
	logger.Info.Printf("Received message: %+v", req)
}
