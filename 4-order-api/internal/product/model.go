package product

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string
	Description string
	Images      pq.StringArray `gorm:"type:text[]"`
}

func NewProduct(name string, description string, images []string) *Product {
	return &Product{
		Name:        name,
		Description: description,
		Images:      images,
	}
}
