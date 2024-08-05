package order_item

import (
	"enterkomputer/internal/domain"
	"gorm.io/gorm"
)

type mysqlOrderItemRepository struct {
	db *gorm.DB
}

func NewMysqlOrderItemRepository(db *gorm.DB) domain.OrderItemRepository {
	return &mysqlOrderItemRepository{db}
}

func (m *mysqlOrderItemRepository) CreateOrderItem(request *domain.OrderItem) error {
	return m.db.Create(request).Error
}

func (m *mysqlOrderItemRepository) GetByOrderID(orderID int) ([]domain.OrderItem, error) {
	var orderItems []domain.OrderItem
	err := m.db.Where("order_id = ?", orderID).Find(&orderItems).Error
	return orderItems, err
}
