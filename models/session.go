package models

import "time"

type Session struct {
	UserId    int64     `json:"userId" db:"userId"`
	SessionId int64     `json:"sessionId" db:"sessionId"`
	CreatedAt time.Time `json:"createdAt" db:"createdAt"`
}
