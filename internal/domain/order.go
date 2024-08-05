package domain

import "time"

type Order struct {
	ID          int       `json:"id" gorm:"primary_key;auto_increment"`
	TableNumber string    `json:"table_number" gorm:"not null; type:varchar(50)"`
	TotalPrice  int       `json:"total_price" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type RequestCreateOrder struct {
	TableNumber string                   `json:"table_number"`
	Items       []RequestCreateOrderItem `json:"items"`
}

type InputCreateOrder struct {
	TableNumber string `json:"table_number"`
	TotalPrice  int    `json:"total_price"`
}

type ResponseCreateOrder struct {
	OrderID     int                       `json:"order_id"`
	TableNumber string                    `json:"table_number"`
	Items       []ResponseCreateOrderItem `json:"items"`
	TotalPrice  int                       `json:"total_price"`
}

func (m *Order) TableName() string {
	return "orders"
}

type OrderRepository interface {
	CreateOrder(request *Order) (*Order, error)
	GetOrder(orderId int) (*Order, error)
}

type OrderService interface {
	CreateOrder(request *RequestCreateOrder) (*ResponseCreateOrder, error)
	GetBill(orderId int) (*ResponseCreateOrder, error)
}
