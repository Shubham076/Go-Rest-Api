package handlers

import "BootCampT1/external/ses"

type SendMailRequest struct {
	EmailConf ses.EmailConfig `json:"email_conf" validate:"required"`
}

type DeleteSessionRequest struct {
	SessionId int `json:"sessionId" validate:"required"`
}
