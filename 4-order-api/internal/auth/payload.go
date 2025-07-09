package auth

type LoginRequest struct {
	Phone string `json:"phone" validate:"required,min=10,max=15"`
}

type LoginResponse struct {
	SessionID string `json:"sessionId"`
}

type VerifyRequest struct {
	Code      int `json:"code" validate:"required,len=4"`
	SessionID string `json:"sessionId" validate:"required,len=32"`
}

type VerifyResponse struct {
	Token string `json:"token"`
}
