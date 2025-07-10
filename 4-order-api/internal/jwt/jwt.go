package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type PhoneNumberKeyType string
type UserIDKeyType string

const (
	PhoneNumberKey PhoneNumberKeyType = "ph"
	UserIDKey      UserIDKeyType      = "sub"
)

type JWTData struct {
	UserID uint
	Phone  string
}

type JWT struct {
	Secret string
}

func NewSecret(secret string) *JWT {
	return &JWT{Secret: secret}
}

func (j *JWT) Create(data JWTData) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ph":  data.Phone,
		"sub": data.UserID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"iat": time.Now().Unix(),
	})

	return t.SignedString([]byte(j.Secret))
}

func (j *JWT) Parse(token string) (bool, *JWTData) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.Secret), nil
	})
	if err != nil {
		return false, nil
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return false, nil
	}

	phone, ok := claims[string(PhoneNumberKey)].(string)
	if !ok {
		return false, nil
	}

	userIDFloat, ok := claims[string(UserIDKey)].(float64)
	if !ok {
		return false, nil
	}

	return parsedToken.Valid, &JWTData{Phone: phone, UserID: uint(userIDFloat)}
}
