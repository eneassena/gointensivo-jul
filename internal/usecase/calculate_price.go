package usecase

import "github.com/eneassena/gointensivo-jul/internal/entity"

type OrderInput struct {
	ID    string
	Price float64
	Tax   float64
}

func NewOrderInput(id string, price, tax float64) *OrderInput {
	return &OrderInput{
		ID:    id,
		Price: price,
		Tax:   tax,
	}
}

type OrderOutPut struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

type calculateFinalPrice struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewCalculateFinalPrice(orderRepository entity.OrderRepositoryInterface) *calculateFinalPrice {
	return &calculateFinalPrice{
		OrderRepository: orderRepository,
	}
}

func (c *calculateFinalPrice) Execute(input OrderInput) (*OrderOutPut, error) {
	order, err := entity.NewOrder(input.ID, input.Price, input.Tax)
	if err != nil {
		return nil, err
	}

	err = order.CalculateFinalPrice()
	if err != nil {
		return nil, err
	}

	err = c.OrderRepository.Save(order)
	if err != nil {
		return nil, err
	}

	return &OrderOutPut{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}, nil
}
