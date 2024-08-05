package order

import (
	"enterkomputer/internal/domain"
	"enterkomputer/internal/middleware/validation"
	"enterkomputer/internal/utilities"
	"github.com/gofiber/fiber/v2"
)

type HttpOrderHandler struct {
	OrderService domain.OrderService
}

func NewHttpHandler(r fiber.Router, orderService domain.OrderService) {
	handler := &HttpOrderHandler{orderService}
	r.Post("/", validation.New[domain.RequestCreateOrder](), handler.CreateOrder)
	r.Get("/:orderId", handler.GetBill)
}

// CreateOrder
//
//	@Summary		Create order
//	@Description	Create order
//	@Tags			Order
//	@Accept			json
//	@Produce		json
//	@Param			request	body		domain.RequestCreateOrder	true	"request body"
//	@Success		200		{object}	domain.ResponseDefault{data=domain.ResponseCreateOrder}
//	@Failure		400		{object}	domain.ResponseDefault{data=nil}
//	@Router			/order [post]
func (h *HttpOrderHandler) CreateOrder(c *fiber.Ctx) error {
	reqCreateOrder := utilities.ExtractStructFromValidator[domain.RequestCreateOrder](c)

	data, err := h.OrderService.CreateOrder(reqCreateOrder)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(domain.ResponseDefault{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Error:   true,
			Data:    nil,
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(domain.ResponseDefault{
		Code:    fiber.StatusOK,
		Message: "Success create order",
		Error:   false,
		Data:    data,
	})
}

// GetBill
//
//	@Summary		Get bill
//	@Description	Get bill
//	@Tags			Order
//	@Accept			json
//	@Produce		json
//	@Param			orderId	path		int	true	"Order ID"
//	@Success		200		{object}	domain.ResponseDefault{data=domain.ResponseCreateOrder}
//	@Failure		400		{object}	domain.ResponseDefault{data=nil}
//	@Router			/order/{orderId} [get]
func (h *HttpOrderHandler) GetBill(c *fiber.Ctx) error {
	orderId, err := c.ParamsInt("orderId")

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(domain.ResponseDefault{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Error:   true,
			Data:    nil,
		})
	}

	data, err := h.OrderService.GetBill(orderId)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(domain.ResponseDefault{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Error:   true,
			Data:    nil,
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(domain.ResponseDefault{
		Code:    fiber.StatusOK,
		Message: "Success get bill",
		Error:   false,
		Data:    data,
	})
}
