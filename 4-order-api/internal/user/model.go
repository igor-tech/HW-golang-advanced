package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	PhoneNumber string `json:"phone_number" validate:"required,min=10,max=15" gorm:"uniqueIndex;size:15"`
	SessionID   string `json:"-" validate:"len=32" gorm:"size:32"`
	Code        string `json:"-" validate:"len=64" gorm:"size:64"`
}

func NewUser(phoneNumber string, sessionID string, code string) *User {
	return &User{
		PhoneNumber: phoneNumber,
		SessionID:   sessionID,
		Code:        code,
	}
}
