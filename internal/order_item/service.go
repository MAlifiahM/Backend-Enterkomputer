package order_item

import "enterkomputer/internal/domain"

type orderItemService struct {
	orderItemRepo domain.OrderItemRepository
}

func NewOrderItemService(orderItemRepo domain.OrderItemRepository) domain.OrderItemService {
	return &orderItemService{
		orderItemRepo,
	}
}
