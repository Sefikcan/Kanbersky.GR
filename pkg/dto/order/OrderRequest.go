package order

type OrderRequest struct {
	CurrencyId int `json:"currency_id"`
	OrderDescription string `json:"order_description"`
	OrderItems []OrderItemRequest `json:"order_items"`
}

type OrderItemRequest struct {
	quantity int `json:"quantity"`
}
