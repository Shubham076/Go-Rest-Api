package models

import "time"

type Otp struct {
	UserId    int64     `json:"userId" db:"userId"`
	Otp       int       `json:"otp" db:"otp"`
	CreatedAt time.Time `json:"createdAt" db:"createdAt"`
}
