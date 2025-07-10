package order

import (
	"order/api/internal/model"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Create(order *model.Order) (*model.Order, error) {

	if err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Omit("Products.*").Create(order).Error; err != nil {
			return err
		}

		return tx.Model(order).
			Association("Products").
			Append(order.Products)
	}); err != nil {
		return nil, err
	}

	if err := r.db.Preload("Products").First(order, order.ID).Error; err != nil {
		return nil, err
	}

	return order, nil
}

func (r *OrderRepository) GetById(id uint) (*model.Order, error) {
	var order model.Order
	err := r.db.
		Model(&model.Order{}).
		Preload("Products").
		Where("id = ?", id).
		First(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepository) GetByUserId(userId uint) ([]model.Order, error) {
	var orders []model.Order
	err := r.db.
		Preload("User").
		Preload("Products").
		Where("user_id = ?", userId).
		Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}
