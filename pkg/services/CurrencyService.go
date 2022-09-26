package services

import (
	dto "GolangRelational/pkg/dto/currency"
	"GolangRelational/pkg/instrastructure"
	"GolangRelational/pkg/instrastructure/entities"
	"GolangRelational/pkg/internal/utils"
)

type CurrencyService interface {
	GetById(id int) (dto.CurrencyResponse, error)
	Add(currency entities.Currency) (dto.CurrencyResponse, error)
	Delete(id int) error
	Update(currency entities.Currency) (dto.CurrencyResponse, error)
	GetAll(pagination utils.Pagination) (*utils.Pagination, error)
}

type currencyService struct {
	currencyRepository instrastructure.CurrencyRepository
}

func (c currencyService) GetById(id int) (dto.CurrencyResponse, error) {
	response, err := c.currencyRepository.GetById(id)
	return dto.MapDto(response), err
}

func (c currencyService) Add(currency entities.Currency) (dto.CurrencyResponse, error) {
	response, err := c.currencyRepository.Add(currency)
	return dto.MapDto(response), err
}

func (c currencyService) Delete(id int) error {
	return c.currencyRepository.Delete(id)
}

func (c currencyService) Update(currency entities.Currency) (dto.CurrencyResponse, error) {
	currentCurrency,err := c.GetById(currency.ID)
	if err != nil {
		return dto.CurrencyResponse{}, err
	}

	currency.IsoCode = currentCurrency.IsoCode
	currency.Title = currentCurrency.Title

	response, err := c.currencyRepository.Update(currency)
	return dto.MapDto(response), err
}

func (c currencyService) GetAll(pagination utils.Pagination) (*utils.Pagination, error) {
	return c.currencyRepository.GetAll(pagination)
}

func NewCurrencyService(c instrastructure.CurrencyRepository) CurrencyService {
	return &currencyService{c}
}