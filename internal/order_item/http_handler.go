package order_item

import (
	"enterkomputer/internal/domain"
	"github.com/gofiber/fiber/v2"
)

type HttpOrderItemHandler struct {
	OrderItemService domain.OrderService
}

func NewHttpHandler(r fiber.Router, orderItemService domain.OrderItemService) {
	//handler := &HttpOrderItemHandler{orderItemService}
	//r.Post("/", validation.New[domain.RequestCreateOrder](), handler.CreateOrder)
	//r.Get(("/:order_id"), handler.GetOrderItems)
}
