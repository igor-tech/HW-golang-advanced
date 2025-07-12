package model

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID uint  `validate:"required" gorm:"not null;constraint:OnDelete:CASCADE"`
	User   *User `json:"user,omitempty" gorm:"foreignKey:UserID"`

	Products []Product `gorm:"many2many:order_products"`
}
