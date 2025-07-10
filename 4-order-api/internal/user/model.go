package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Phone     string `json:"phone" validate:"required,min=10,max=15" gorm:"uniqueIndex;size:15"`
	SessionID string `json:"-" validate:"len=32" gorm:"size:32"`
	Code      string `json:"-" validate:"len=64" gorm:"size:64"`
}

func NewUser(phone string, sessionID string, code string) *User {
	return &User{
		Phone:     phone,
		SessionID: sessionID,
		Code:      code,
	}
}
