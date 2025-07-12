package order

import (
	"errors"
	"order/api/internal/model"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Create(order *model.Order, productsIDs []uint) (*model.Order, error) {
	if len(productsIDs) == 0 {
		return nil, errors.New("productsIDs is empty")
	}

	// Remove duplicates
	uniq := make(map[uint]struct{}, len(productsIDs))
	for _, id := range productsIDs {
		uniq[id] = struct{}{}
	}
	dedup := make([]uint, 0, len(uniq))
	for id := range uniq {
		dedup = append(dedup, id)
	}

	if err := r.db.Transaction(func(tx *gorm.DB) error {
		// Check that all products exist
		if err := r.ValidateProductExist(tx, dedup); err != nil {
			return err
		}

		// Create order first
		if err := tx.Create(order).Error; err != nil {
			return err
		}

		// Load existing products from database
		var products []model.Product
		if err := tx.Where("id IN (?)", dedup).Find(&products).Error; err != nil {
			return err
		}

		// Associate products with order
		if err := tx.Model(order).Association("Products").Append(products); err != nil {
			return err
		}

		return tx.Preload("Products").First(order, order.ID).Error
	}); err != nil {
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

func (r *OrderRepository) ValidateProductExist(tx *gorm.DB, ids []uint) error {
	var cnt int64

	err := tx.Model(&model.Product{}).Where("id IN (?)", ids).Count(&cnt).Error
	if err != nil {
		return err
	}

	if cnt != int64(len(ids)) {
		return errors.New("some products not found")
	}
	return nil
}
