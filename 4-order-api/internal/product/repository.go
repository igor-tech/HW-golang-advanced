package product

import (
	"order/api/internal/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetAll() ([]model.Product, error) {
	var products []model.Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) Create(product *model.Product) (*model.Product, error) {
	if err := r.db.Create(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (r *ProductRepository) Delete(id uint) error {
	if err := r.db.Delete(&model.Product{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *ProductRepository) Update(product *model.Product) (*model.Product, error) {
	if err := r.db.Clauses(clause.Returning{}).Updates(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (r *ProductRepository) GetById(id uint) (*model.Product, error) {
	var product model.Product
	if err := r.db.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}
