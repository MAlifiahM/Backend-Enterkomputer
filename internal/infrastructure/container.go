package infrastructure

import (
	"enterkomputer/internal/config"
	"enterkomputer/internal/domain"
	"enterkomputer/internal/order"
	"enterkomputer/internal/order_item"
	"enterkomputer/internal/product"
	"enterkomputer/pkg/xlogger"
	"github.com/caarlos0/env/v11"
)

var (
	cfg config.Config

	ProductRepository domain.ProductRepository
	ProductService    domain.ProductService

	OrderRepository domain.OrderRepository
	OrderService    domain.OrderService

	OrderItemRepository domain.OrderItemRepository
	OrderItemService    domain.OrderItemService
)

func init() {
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}

	xlogger.Setup(cfg)
	dbSetup()

	ProductRepository = product.NewMysqlProductRepository(db)
	ProductService = product.NewProductService(ProductRepository)

	OrderItemRepository = order_item.NewMysqlOrderItemRepository(db)
	OrderItemService = order_item.NewOrderItemService(OrderItemRepository)

	OrderRepository = order.NewMysqlOrderRepository(db)
	OrderService = order.NewOrderService(OrderRepository, ProductRepository, OrderItemRepository)
}
