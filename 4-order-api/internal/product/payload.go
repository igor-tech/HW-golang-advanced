package product

import "github.com/lib/pq"

type CreateProductRequest struct {
	Name        string         `json:"name" validate:"required,min=3,max=255"`
	Description string         `json:"description" validate:"required,min=3,max=255"`
	Images      pq.StringArray `json:"images" validate:"required,min=1,max=10"`
}

type UpdateProductRequest struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Images      pq.StringArray `json:"images"`
}
