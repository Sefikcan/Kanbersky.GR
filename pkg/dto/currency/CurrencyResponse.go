package currency

import (
	"GolangRelational/pkg/instrastructure/entities"
)

type CurrencyResponse struct {
	ID int `json:"id"`
	Title string `json:"title"`
	IsoCode string `json:"iso_code"`
}

type CurrencyResponses []*CurrencyResponse

func MapDto(currency entities.Currency) CurrencyResponse {
	return CurrencyResponse{
		ID:      currency.ID,
		Title:   currency.Title,
		IsoCode: currency.IsoCode,
	}
}

func MapListDto(currencies []*entities.Currency) CurrencyResponses {
	currencyResp := CurrencyResponses{}
	for _, c := range currencies {
		currency := MapDto(*c)
		currencyResp = append(currencyResp, &currency)
	}
	return currencyResp
}

