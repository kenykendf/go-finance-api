package service

import (
	"errors"

	"github.com/go-finance/internal/app/model"
	"github.com/go-finance/internal/app/schema"
)

type CurrencyService struct {
	currency CurrencyRepo
}

type CurrencyRepo interface {
	Create(params *model.Currency) error
	GetCurrenciesLists() ([]model.Currency, error)
}

func NewCurrencyService(currency CurrencyRepo) *CurrencyService {
	return &CurrencyService{currency: currency}
}

func (cs *CurrencyService) Create(params *schema.CreateCurrency) error {
	var data model.Currency

	data.Country = params.Country
	data.Currency = params.Currency
	data.CurrencyAbb = params.CurrencyAbb

	err := cs.currency.Create(&data)
	if err != nil {
		return errors.New("unable to create new currency")
	}
	return nil
}

func (cs *CurrencyService) GetCurrenciesLists() ([]schema.GetCurrencyLists, error) {
	var response []schema.GetCurrencyLists

	currencies, err := cs.currency.GetCurrenciesLists()
	if err != nil {
		return nil, errors.New("unable to get currencies lists")
	}

	for _, v := range currencies {
		var data schema.GetCurrencyLists
		data.ID = v.ID
		data.Country = v.Country
		data.Currency = v.Currency
		data.CurrencyAbb = v.CurrencyAbb

		response = append(response, data)
	}

	return response, nil
}
