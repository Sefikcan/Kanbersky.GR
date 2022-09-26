package instrastructure

import (
	dto "GolangRelational/pkg/dto/order"
	"GolangRelational/pkg/instrastructure/entities"
	"GolangRelational/pkg/internal/utils"
	"gorm.io/gorm"
)

type OrderRepository interface {
	GetById(id int) (entities.Order, error)
	Add(order entities.Order) (entities.Order, error)
	Delete(id int) error
	Update(order entities.Order) (entities.Order, error)
	GetAll(pagination utils.Pagination) (*utils.Pagination, error)
}

type orderRepository struct {
	db *gorm.DB
}

func (o orderRepository) GetById(id int) (entities.Order, error) {
	order := entities.Order{}
	err := o.db.
		Preload("Currency").
		Preload("OrderItems").
		First(&order, `orders.id = ?`, id).Error
	return order, err
}

func (o orderRepository) Add(order entities.Order) (entities.Order, error) {
	if result := o.db.Save(&order); result.Error != nil {
		return order, result.Error
	}

	return order, nil
}

func (o orderRepository) Delete(id int) error {
	if result := o.db.Select("OrderItems").Delete(&entities.Order{Id: id}); result.Error != nil {
		return result.Error
	}

	return nil
}

func (o orderRepository) Update(order entities.Order) (entities.Order, error) {
	if result := o.db.Save(&order); result.Error != nil {
		return order, result.Error
	}

	return order, nil
}

func (o orderRepository) GetAll(pagination utils.Pagination) (*utils.Pagination, error) {
	var orders []*entities.Order
	o.db.Preload("Currency").Preload("OrderItems").Scopes(utils.Paginate(orders, &pagination, o.db)).Find(&orders)
	pagination.Rows = dto.MapListDto(orders)


	return &pagination, nil
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db:db}
}
