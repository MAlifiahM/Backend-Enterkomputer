package order

import (
	"enterkomputer/internal/domain"
	"gorm.io/gorm"
)

type mysqlOrderRepository struct {
	db *gorm.DB
}

func NewMysqlOrderRepository(db *gorm.DB) domain.OrderRepository {
	return &mysqlOrderRepository{db}
}

func (m *mysqlOrderRepository) CreateOrder(request *domain.Order) (*domain.Order, error) {
	err := m.db.Create(request).Error

	if err != nil {
		return nil, err
	}

	return request, nil
}

func (m *mysqlOrderRepository) GetOrder(orderId int) (*domain.Order, error) {
	var order domain.Order
	err := m.db.Where("id = ?", orderId).Find(&order).Error
	return &order, err
}
