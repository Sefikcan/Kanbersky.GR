package services

import (
	dto "GolangRelational/pkg/dto/order"
	"GolangRelational/pkg/instrastructure"
	"GolangRelational/pkg/instrastructure/entities"
	"GolangRelational/pkg/internal/utils"
)

type OrderService interface {
	GetById(id int) (dto.OrderResponse, error)
	Add(order entities.Order) (dto.OrderResponse, error)
	Delete(id int) error
	Update(order entities.Order) (dto.OrderResponse, error)
	GetAll(pagination utils.Pagination) (*utils.Pagination, error)
}

type orderService struct {
	orderRepository instrastructure.OrderRepository
	currencyService CurrencyService
}

func (o orderService) GetById(id int) (dto.OrderResponse, error) {
	response, err := o.orderRepository.GetById(id)
	return dto.MapDto(response), err
}

func (o orderService) Add(order entities.Order) (dto.OrderResponse, error) {
	orderCurrency, err := o.currencyService.GetById(order.CurrencyId)
	if err != nil {
		return dto.OrderResponse{},err
	}
	order.Currency = entities.Currency{
		ID: orderCurrency.ID,
		Title: orderCurrency.Title,
		IsoCode: orderCurrency.IsoCode,
	}
	response, err := o.orderRepository.Add(order)
	return dto.MapDto(response), err
}

func (o orderService) Delete(id int) error {
	return o.orderRepository.Delete(id)
}

func (o orderService) Update(order entities.Order) (dto.OrderResponse, error) {
	currentOrder,err := o.GetById(order.Id)
	if err != nil {
		return dto.OrderResponse{}, err
	}

	currencyId := currentOrder.Currency.ID
	if order.CurrencyId != 0 {
		currencyId = order.CurrencyId
	}

	orderCurrency, err := o.currencyService.GetById(currencyId)
	if err != nil {
		return dto.OrderResponse{},err
	}
	order.Currency = entities.Currency{
		ID: orderCurrency.ID,
		Title: orderCurrency.Title,
		IsoCode: orderCurrency.IsoCode,
	}

	if order.OrderDescription == "" {
		order.OrderDescription = currentOrder.OrderDescription
	}

	if len(order.OrderItems) > 0 {
		order.OrderItems = append(order.OrderItems)
	}

	for i := 0; i < len(currentOrder.OrderItemResponse); i++ {
		order.OrderItems = append(order.OrderItems, dto.MapEntity(currentOrder.OrderItemResponse[i]))
	}

	response, err := o.orderRepository.Update(order)
	return dto.MapDto(response), err
}

func (o orderService) GetAll(pagination utils.Pagination) (*utils.Pagination, error) {
	return o.orderRepository.GetAll(pagination)
}

func NewOrderService(orderRepository instrastructure.OrderRepository, currencyService CurrencyService) OrderService {
	return &orderService{
		orderRepository: orderRepository,
		currencyService: currencyService,
	}
}
