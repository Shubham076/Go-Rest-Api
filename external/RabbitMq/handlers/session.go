package handlers

import (
	"BootCampT1/external/rds/queries"
	"BootCampT1/logger"
	"encoding/json"
	"github.com/streadway/amqp"
)

func DeleteSessionHandler(msg amqp.Delivery) {
	var req DeleteSessionRequest
	err := json.Unmarshal(msg.Body, &req)

	if err != nil {
		logger.Error.Printf("Unable to unmarshall msg due to err: %s", err.Error())
		return
	}

	_, err = queries.DeleteSession(req.SessionId)
	if err != nil {
		logger.Error.Println("Unable to delete user session, err: %v", err)
		return
	}
}
