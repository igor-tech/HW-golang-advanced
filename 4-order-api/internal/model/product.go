package model

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string         `json:"name" validate:"required,min=3,max=255"`
	Description string         `json:"description" validate:"required,min=3,max=255"`
	Images      pq.StringArray `json:"images" validate:"required,min=1,max=10" gorm:"type:text[]" `

	Orders []Order `gorm:"many2many:order_products"`
}

func NewProduct(name string, description string, images []string) *Product {
	return &Product{
		Name:        name,
		Description: description,
		Images:      images,
	}
}
