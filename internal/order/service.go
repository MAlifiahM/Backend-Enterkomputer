package order

import (
	"enterkomputer/internal/domain"
)

type orderService struct {
	orderRepo     domain.OrderRepository
	productRepo   domain.ProductRepository
	orderItemRepo domain.OrderItemRepository
}

func NewOrderService(orderRepo domain.OrderRepository, productRepo domain.ProductRepository, orderItemRepo domain.OrderItemRepository) domain.OrderService {
	return &orderService{
		orderRepo,
		productRepo,
		orderItemRepo,
	}
}

func (o *orderService) CreateOrder(request *domain.RequestCreateOrder) (*domain.ResponseCreateOrder, error) {
	itemData := make([]domain.ProductResponse, 0)
	allTotalPrice := 0

	for _, item := range request.Items {
		dataItem, err := o.productRepo.ProductByID(item.ProductID)

		if err != nil {
			return nil, err
		}

		itemData = append(itemData, domain.ProductResponse{
			ID:         dataItem.ID,
			Category:   dataItem.Category,
			Quantity:   item.Quantity,
			Name:       dataItem.Name,
			Variant:    dataItem.Variant,
			Price:      dataItem.Price,
			TotalPrice: dataItem.Price * item.Quantity,
		})

		allTotalPrice += dataItem.Price * item.Quantity
	}

	dataOrder, err := o.orderRepo.CreateOrder(&domain.Order{
		TableNumber: request.TableNumber,
		TotalPrice:  allTotalPrice,
	})

	if err != nil {
		return nil, err
	}

	var responseData domain.ResponseCreateOrder
	var ItemResponseData []domain.ResponseCreateOrderItem
	ItemResponseData = append(ItemResponseData, domain.ResponseCreateOrderItem{
		Printer: "A",
	})
	ItemResponseData = append(ItemResponseData, domain.ResponseCreateOrderItem{
		Printer: "B",
	})
	ItemResponseData = append(ItemResponseData, domain.ResponseCreateOrderItem{
		Printer: "C",
	})
	for _, item := range itemData {
		reqOrderItem := &domain.OrderItem{
			OrderID:   dataOrder.ID,
			ProductID: item.ID,
			Quantity:  item.Quantity,
			Price:     item.TotalPrice,
		}

		err = o.orderItemRepo.CreateOrderItem(reqOrderItem)
		if err != nil {
			return nil, err
		}

		if item.Category == "makanan" {
			ItemResponseData[1].Product = append(ItemResponseData[1].Product, domain.ResponseItemsCreateOrderItem{
				Category:   item.Category,
				Name:       item.Name,
				Variant:    item.Variant,
				Quantity:   item.Quantity,
				Price:      item.Price,
				TotalPrice: item.TotalPrice,
			})
			ItemResponseData[0].Product = append(ItemResponseData[0].Product, domain.ResponseItemsCreateOrderItem{
				Category:   item.Category,
				Name:       item.Name,
				Variant:    item.Variant,
				Quantity:   item.Quantity,
				Price:      item.Price,
				TotalPrice: item.TotalPrice,
			})
		} else if item.Category == "minuman" {
			ItemResponseData[2].Product = append(ItemResponseData[2].Product, domain.ResponseItemsCreateOrderItem{
				Category:   item.Category,
				Name:       item.Name,
				Variant:    item.Variant,
				Quantity:   item.Quantity,
				Price:      item.Price,
				TotalPrice: item.TotalPrice,
			})
			ItemResponseData[0].Product = append(ItemResponseData[0].Product, domain.ResponseItemsCreateOrderItem{
				Category:   item.Category,
				Name:       item.Name,
				Variant:    item.Variant,
				Quantity:   item.Quantity,
				Price:      item.Price,
				TotalPrice: item.TotalPrice,
			})
		} else if item.Category == "promo" {
			ItemResponseData[1].Product = append(ItemResponseData[1].Product, domain.ResponseItemsCreateOrderItem{
				Category:   item.Category,
				Name:       item.Name,
				Variant:    item.Variant,
				Quantity:   item.Quantity,
				Price:      item.Price,
				TotalPrice: item.TotalPrice,
			})
			ItemResponseData[2].Product = append(ItemResponseData[2].Product, domain.ResponseItemsCreateOrderItem{
				Category:   item.Category,
				Name:       item.Name,
				Variant:    item.Variant,
				Quantity:   item.Quantity,
				Price:      item.Price,
				TotalPrice: item.TotalPrice,
			})
			ItemResponseData[0].Product = append(ItemResponseData[0].Product, domain.ResponseItemsCreateOrderItem{
				Category:   item.Category,
				Name:       item.Name,
				Variant:    item.Variant,
				Quantity:   item.Quantity,
				Price:      item.Price,
				TotalPrice: item.TotalPrice,
			})
		}
	}

	responseData.TableNumber = dataOrder.TableNumber
	responseData.TotalPrice = dataOrder.TotalPrice
	responseData.OrderID = dataOrder.ID
	responseData.Items = ItemResponseData

	return &responseData, nil
}

func (o *orderService) GetBill(id int) (*domain.ResponseCreateOrder, error) {
	order, err := o.orderRepo.GetOrder(id)

	if err != nil {
		return nil, err
	}

	orderItem, err := o.orderItemRepo.GetByOrderID(order.ID)

	if err != nil {
		return nil, err
	}

	var ItemResponseData []domain.ResponseCreateOrderItem

	ItemResponseData = append(ItemResponseData, domain.ResponseCreateOrderItem{
		Printer: "A",
	})
	for _, item := range orderItem {
		dataItem, err := o.productRepo.ProductByID(item.ProductID)

		if err != nil {
			return nil, err
		}
		ItemResponseData[0].Product = append(ItemResponseData[0].Product, domain.ResponseItemsCreateOrderItem{
			Category:   dataItem.Category,
			Name:       dataItem.Name,
			Variant:    dataItem.Variant,
			Quantity:   item.Quantity,
			Price:      dataItem.Price,
			TotalPrice: item.Price,
		})
	}

	responseData := domain.ResponseCreateOrder{
		OrderID:     order.ID,
		TableNumber: order.TableNumber,
		TotalPrice:  order.TotalPrice,
		Items:       ItemResponseData,
	}

	return &responseData, nil
}
