package domain

import "time"

type Product struct {
	ID        int       `json:"id" gorm:"primary_key;auto_increment"`
	Category  string    `json:"category" gorm:"not null; type:varchar(50)"`
	Name      string    `json:"name" gorm:"not null; type:varchar(50)"`
	Variant   string    `json:"variant,omitempty" gorm:"type:varchar(50)"`
	Price     int       `json:"price" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RequestCreateOrderItem struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type ResponseItemsCreateOrderItem struct {
	Category   string `json:"category"`
	Name       string `json:"name"`
	Variant    string `json:"variant"`
	Quantity   int    `json:"quantity"`
	Price      int    `json:"price"`
	TotalPrice int    `json:"total_price"`
}

type ResponseCreateOrderItem struct {
	Printer string                         `json:"printer"`
	Product []ResponseItemsCreateOrderItem `json:"product"`
}

type ProductResponse struct {
	ID         int    `json:"id"`
	Category   string `json:"category"`
	Name       string `json:"name"`
	Variant    string `json:"variant,omitempty"`
	Price      int    `json:"price"`
	TotalPrice int    `json:"total_price"`
	Quantity   int    `json:"quantity"`
}

func (m *Product) TableName() string {
	return "products"
}

type ProductRepository interface {
	GetAllProducts() ([]Product, error)
	ProductByID(id int) (*Product, error)
}

type ProductService interface {
	GetAllProducts() ([]Product, error)
}
