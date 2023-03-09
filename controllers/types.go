package controllers

type LoginRequest struct {
	Email string `json:"email" validate:"required"`
}
type SignUpRequest struct {
	Email    string `json:"email" validate:"required"`
	Username string `json:"username" validate:"required"`
}
type LogoutRequest struct {
	UserId    int64 `json:"userId" validate:"required"`
	SessionId int64 `json:"sessionId" validate:"required"`
}

type GetUserDetailsRequest struct {
	Email string `json:"email" validate:"required"`
}

type VerifyOtpRequest struct {
	Otp    int   `json:"otp"       validate:"required"`
	UserId int64 `json:"userId"    validate:"required"`
}
