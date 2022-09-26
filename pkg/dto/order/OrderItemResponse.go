package order

import (
	"GolangRelational/pkg/instrastructure/entities"
)

func MapListEntity(orderItemResponses []OrderItemResponse) []entities.OrderItem{
	var orderItems []entities.OrderItem
	for _,item := range orderItemResponses {
		var orderItem entities.OrderItem
		orderItem.Quantity = item.Quantity
		orderItem.Id = item.Id
		orderItems = append(orderItems, orderItem)
	}

	return orderItems
}

func MapEntity(response OrderItemResponse) entities.OrderItem {
	return entities.OrderItem{
		Id: response.Id,
		Quantity: response.Quantity,
	}
}
