package domain

import "time"

type Printer struct {
	ID          int       `json:"id" gorm:"primary_key;auto_increment"`
	Name        string    `json:"name" gorm:"not null; type:varchar(50)"`
	Description string    `json:"description" gorm:"type:varchar(50)"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type PrinterRepository interface {
}

type PrinterService interface {
}
