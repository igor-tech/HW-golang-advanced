package middleware

import (
	"context"
	"net/http"
	"order/api/configs"
	"order/api/internal/jwt"
	"strings"
)

func IsAuth(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			http.Error(w, "Invalid Authorization format", http.StatusUnauthorized)
			return
		}

		token := strings.TrimSpace(parts[1])
		isValid, data := jwt.NewSecret(config.JwtSecret).Parse(token)
		if !isValid {
			http.Error(w, "Token expired", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), jwt.PhoneNumberKey, data.Phone)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
