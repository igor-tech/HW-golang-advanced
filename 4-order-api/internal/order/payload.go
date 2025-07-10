package order

import (
	"time"

	"github.com/lib/pq"
)

type CreateOrderRequest struct {
	ProductIDs []uint `json:"product_ids" validate:"required,min=1"`
}

type CreateOrderResponse struct {
	OrderID    uint      `json:"order_id"`
	ProductsID []uint    `json:"products_id"`
	UserID     uint      `json:"user_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Product struct {
	ID     uint           `json:"id"`
	Name   string         `json:"name"`
	Images pq.StringArray `json:"images"`
}

type GetOrderResponse struct {
	OrderID   uint      `json:"order_id"`
	Products  []Product `json:"products"`
	CreatedAt time.Time `json:"created_at"`
}
