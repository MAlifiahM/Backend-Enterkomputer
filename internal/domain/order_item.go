package domain

import "time"

type OrderItem struct {
	ID        int       `json:"id" gorm:"primary_key;auto_increment"`
	OrderID   int       `json:"order_id" gorm:"not null"`
	ProductID int       `json:"product_id" gorm:"not null"`
	Quantity  int       `json:"quantity" gorm:"not null"`
	Price     int       `json:"price" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *OrderItem) TableName() string {
	return "order_items"
}

type OrderItemRepository interface {
	CreateOrderItem(request *OrderItem) error
	GetByOrderID(orderID int) ([]OrderItem, error)
}

type OrderItemService interface {
}
