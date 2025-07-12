package order

import (
	"order/api/internal/model"

	"gorm.io/gorm"
)

func idsToProduct(ids []uint) []model.Product {
	products := make([]model.Product, len(ids))
	for i, id := range ids {
		products[i] = model.Product{
			Model: gorm.Model{
				ID: id,
			},
		}
	}
	return products
}

func orderToResponse(order *model.Order) CreateOrderResponse {
	if order == nil {
		return CreateOrderResponse{}
	}

	productsIDs := make([]uint, len(order.Products))
	for i, product := range order.Products {
		productsIDs[i] = product.ID
	}

	return CreateOrderResponse{
		OrderID:    order.ID,
		ProductsID: productsIDs,
		UserID:     order.UserID,
		CreatedAt:  order.CreatedAt,
		UpdatedAt:  order.UpdatedAt,
	}
}

func toDTO(o model.Order) GetOrderResponse {
	products := make([]Product, len(o.Products))
	for i, p := range o.Products {
		products[i] = Product{ID: p.ID, Name: p.Name, Images: p.Images}
	}
	return GetOrderResponse{
		OrderID:   o.ID,
		Products:  products,
		CreatedAt: o.CreatedAt,
	}
}
