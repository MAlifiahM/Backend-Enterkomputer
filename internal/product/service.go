package product

import "enterkomputer/internal/domain"

type productService struct {
	productRepo domain.ProductRepository
}

func NewProductService(productRepo domain.ProductRepository) domain.ProductService {
	return &productService{
		productRepo,
	}
}

func (a *productService) GetAllProducts() ([]domain.Product, error) {
	return a.productRepo.GetAllProducts()
}
