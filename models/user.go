package models

type User struct {
	Id       int64  `db:"userId"`
	Email    string `json:"email" db:"email" validate:"required"`
	Username string `json:"username" db:"username" validate:"required"`
}
