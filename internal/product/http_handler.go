package product

import (
	"enterkomputer/internal/domain"
	"github.com/gofiber/fiber/v2"
)

type HttpProductHandler struct {
	ProductService domain.ProductService
}

func NewHttpHandler(r fiber.Router, productService domain.ProductService) {
	handler := &HttpProductHandler{productService}
	r.Get("/all", handler.GetAllProducts)
}

// GetAllProducts
//
//	@Summary		Get all products
//	@Description	Get all products
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	domain.ResponseDefault{data=[]domain.Product}
//	@Failure		400	{object}	domain.ResponseDefault{data=nil}
//	@Router			/product/all [get]
func (h *HttpProductHandler) GetAllProducts(c *fiber.Ctx) error {
	product, err := h.ProductService.GetAllProducts()
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(
			domain.ResponseDefault{
				Code:    fiber.StatusBadRequest,
				Message: error.Error(err),
				Error:   true,
				Data:    nil,
			})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(domain.ResponseDefault{
		Code:    fiber.StatusOK,
		Message: "Success get all products",
		Error:   false,
		Data:    product,
	})
}
