package product

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetAll() ([]Product, error) {
	var products []Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) Create(product *Product) (*Product, error) {
	if err := r.db.Create(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (r *ProductRepository) Delete(id uint) error {
	if err := r.db.Delete(&Product{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *ProductRepository) Update(product *Product) (*Product, error) {
	if err := r.db.Clauses(clause.Returning{}).Updates(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (r *ProductRepository) GetById(id uint) (*Product, error) {
	var product Product
	if err := r.db.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}
