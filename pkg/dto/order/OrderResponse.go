package order

import (
	"GolangRelational/pkg/dto/currency"
	"GolangRelational/pkg/instrastructure/entities"
)

type OrderResponse struct {
	Id int `json:"id"`
	OrderDescription string `json:"order_description"`
	Currency currency.CurrencyResponse `json:"currency"`
	OrderItemResponse []OrderItemResponse `json:"order_items"`
}

type OrderItemResponse struct {
	Id int `json:"id"`
	Quantity int `json:"quantity"`
}

type OrderResponses []*OrderResponse

func MapDto(order entities.Order) OrderResponse {
	var orderItemResponses []OrderItemResponse
	for _,item := range order.OrderItems {
		var orderItemResponse OrderItemResponse
		orderItemResponse.Quantity = item.Quantity
		orderItemResponse.Id = item.Id
		orderItemResponses = append(orderItemResponses, orderItemResponse)
	}

	return OrderResponse{
		Id : order.Id,
		OrderDescription: order.OrderDescription,
		Currency: currency.CurrencyResponse{
			ID: order.Currency.ID,
			IsoCode: order.Currency.IsoCode,
			Title: order.Currency.Title,
		},
		OrderItemResponse: orderItemResponses,
	}
}

func MapListDto(orders []*entities.Order) OrderResponses {
	ordersResp := OrderResponses{}
	for _, o := range orders {
		order := MapDto(*o)
		ordersResp = append(ordersResp, &order)
	}
	return ordersResp
}
