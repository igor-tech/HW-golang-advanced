package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type PhoneNumberKeyType string

const PhoneNumberKey PhoneNumberKeyType = "phone_number"

type JWTData struct {
	Phone string
}

type JWT struct {
	Secret string
}

func NewSecret(secret string) *JWT {
	return &JWT{Secret: secret}
}

func (j *JWT) Create(data JWTData) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		string(PhoneNumberKey): data.Phone,
		"exp":                  time.Now().Add(time.Hour * 24).Unix(),
		"iat":                  time.Now().Unix(),
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

	phone, ok := parsedToken.Claims.(jwt.MapClaims)[string(PhoneNumberKey)].(string)
	if !ok {
		return false, nil
	}

	return parsedToken.Valid, &JWTData{Phone: phone}
}
