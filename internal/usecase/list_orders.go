package usecase

import "github.com/devfullcycle/20-CleanArch/internal/entity"

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

type ListOrdersOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

func NewListOrdersUseCase(OrderRepository entity.OrderRepositoryInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *ListOrdersUseCase) Execute() ([]ListOrdersOutputDTO, error) {
	orders, err := c.OrderRepository.FindAll()

	if err != nil {
		return nil, err
	}

	var output []ListOrdersOutputDTO
	for _, order := range orders {
		output = append(output, ListOrdersOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}

	return output, nil
}
