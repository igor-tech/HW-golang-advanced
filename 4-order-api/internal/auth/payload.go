package auth

type LoginRequest struct {
	Phone string `json:"phone" validate:"required,min=10,max=15"`
}

type LoginResponse struct {
	SessionID string `json:"sessionId"`
}

type VerifyRequest struct {
	Code      int    `json:"code" validate:"required,min=1000,max=9999"`
	SessionID string `json:"sessionId" validate:"required,min=32,max=32"`
}

type VerifyResponse struct {
	Token string `json:"token"`
}
