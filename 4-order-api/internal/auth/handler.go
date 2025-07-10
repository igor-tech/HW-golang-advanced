package auth

import (
	"errors"
	"fmt"
	"net/http"
	"order/api/internal/jwt"
	"order/api/internal/model"
	"order/api/internal/user"
	"order/api/pkg/request"
	"order/api/pkg/response"

	"gorm.io/gorm"
)

type AuthHandler struct {
	UserRepository *user.UserRepository
	JWT            *jwt.JWT
}

type AuthHandlerDeps struct {
	UserRepository *user.UserRepository
	JWT            *jwt.JWT
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	authHandler := &AuthHandler{UserRepository: deps.UserRepository, JWT: deps.JWT}

	router.HandleFunc("POST /auth/login", authHandler.Login())
	router.HandleFunc("POST /auth/verify", authHandler.Verify())
}

func (h *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := request.HandleBody[LoginRequest](w, r)
		if err != nil {
			return
		}

		sessionID, err := user.GenerateSessionID()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		code, err := user.GenerateCode()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println(code) // SEND SMS Service
		codeHash := user.HashCode(int(code), sessionID)

		u := model.User{
			Phone:     payload.Phone,
			SessionID: sessionID,
			Code:      codeHash,
		}

		if err := h.UserRepository.UpsertSession(&u); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response.Response(w, http.StatusOK, LoginResponse{SessionID: sessionID})
	}
}

func (h *AuthHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := request.HandleBody[VerifyRequest](w, r)
		if err != nil {
			return
		}

		usr, err := h.UserRepository.FindBySessionId(payload.SessionID)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Invalid session", http.StatusUnauthorized)
			return
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if !user.CheckHash(payload.Code, usr.SessionID, usr.Code) {
			http.Error(w, "Invalid code", http.StatusUnauthorized)
			return
		}

		token, err := h.JWT.Create(jwt.JWTData{Phone: usr.Phone, UserID: usr.ID})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err = h.UserRepository.UpdateFields(usr.ID, map[string]interface{}{"code": "", "session_id": ""}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response.Response(w, http.StatusOK, VerifyResponse{Token: token})
	}
}
