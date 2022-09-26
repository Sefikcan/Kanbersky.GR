package entities

import (
	"time"
)

type OrderItem struct {
	Id int `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Order Order `json:"order"`
	OrderId int `gorm:"column:order_id" json:"order_id"`
	Quantity int `json:"quantity"`
}
