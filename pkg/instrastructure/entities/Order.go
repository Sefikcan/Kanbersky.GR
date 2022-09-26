package entities

import (
	"time"
)

type Order struct {
	Id int `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CurrencyId int `json:"currency_id"`
	Currency Currency `gorm:"foreignKey:currency_id" json:"currency"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderId" json:"order_items"`
	OrderDescription string `json:"order_description"`
}
