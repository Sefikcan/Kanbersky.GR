package instrastructure

import (
	dto "GolangRelational/pkg/dto/currency"
	"GolangRelational/pkg/instrastructure/entities"
	"GolangRelational/pkg/internal/utils"
	"gorm.io/gorm"
)

type CurrencyRepository interface {
	GetById(id int) (entities.Currency, error)
	Add(currency entities.Currency) (entities.Currency, error)
	Delete(id int) error
	Update(currency entities.Currency) (entities.Currency, error)
	GetAll(pagination utils.Pagination) (*utils.Pagination, error)
}

type currencyRepository struct {
	db *gorm.DB
}

func (c currencyRepository) GetById(id int) (entities.Currency, error) {
	currency := entities.Currency{}
	err := c.db.Where(`id = ?`, id).First(&currency).Error
	return currency, err
}

func (c currencyRepository) Add(currency entities.Currency) (entities.Currency, error) {
	if result := c.db.Create(&currency); result.Error != nil {
		return currency, result.Error
	}

	return currency, nil
}

func (c currencyRepository) Delete(id int) error {
	if result := c.db.Delete(&entities.Currency{ID: id}); result.Error != nil {
		return result.Error
	}

	return nil
}

func (c currencyRepository) Update(currency entities.Currency) (entities.Currency, error) {
	if result := c.db.Save(&currency); result.Error != nil {
		return currency, result.Error
	}

	return currency, nil
}

func (c currencyRepository) GetAll(pagination utils.Pagination) (*utils.Pagination, error) {
	var currencies []*entities.Currency
	c.db.Scopes(utils.Paginate(currencies, &pagination, c.db)).Find(&currencies)
	pagination.Rows = dto.MapListDto(currencies)

	return &pagination, nil
}

func NewCurrencyRepository(db *gorm.DB) CurrencyRepository {
	return &currencyRepository{db:db}
}
