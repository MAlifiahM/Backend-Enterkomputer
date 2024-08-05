package product

import (
	"enterkomputer/internal/domain"
	"errors"
	"gorm.io/gorm"
)

type mysqlProductRepository struct {
	db *gorm.DB
}

func NewMysqlProductRepository(db *gorm.DB) domain.ProductRepository {
	return &mysqlProductRepository{db}
}

func (m *mysqlProductRepository) GetAllProducts() ([]domain.Product, error) {
	var products []domain.Product
	err := m.db.Find(&products).Error
	return products, err
}

func (m *mysqlProductRepository) ProductByID(id int) (*domain.Product, error) {
	var product domain.Product
	if err := m.db.Where("id = ?", id).Find(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &product, nil
}
