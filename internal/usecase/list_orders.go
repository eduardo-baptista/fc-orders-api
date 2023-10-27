package usecase

import "github.com/devfullcycle/20-CleanArch/internal/entity"

type OrdersDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type ListOrdersOutputDTO struct {
	Orders []OrdersDTO `json:"orders"`
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(OrderRepository entity.OrderRepositoryInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (l *ListOrdersUseCase) Execute() (ListOrdersOutputDTO, error) {
	orders, err := l.OrderRepository.List()
	if err != nil {
		return ListOrdersOutputDTO{}, err
	}

	var ordersDTO []OrdersDTO
	for _, order := range orders {
		ordersDTO = append(ordersDTO, OrdersDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}

	return ListOrdersOutputDTO{
		Orders: ordersDTO,
	}, nil
}
