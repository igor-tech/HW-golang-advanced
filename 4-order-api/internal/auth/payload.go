package auth

type LoginRequest struct {
	PhoneNumber string `json:"phone_number" validate:"required,min=10,max=15"`
}

type LoginResponse struct {
	SessionID string `json:"session_id"`
}

type VerifyRequest struct {
	Code      string `json:"code" validate:"required,len=4"`
	SessionID string `json:"session_id" validate:"required,len=32"`
}

type VerifyResponse struct {
	Token string `json:"token"`
}
