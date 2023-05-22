package service

import (
	"errors"
	"fmt"

	"github.com/go-finance/internal/app/model"
	"github.com/go-finance/internal/app/schema"

	log "github.com/sirupsen/logrus"
)

type CurrencyService struct {
	currency CurrencyRepo
}

type CurrencyRepo interface {
	CreateCurrency(params *model.Currency) error
	GetCurrenciesLists() ([]model.Currency, error)
	GetCurrencyByID(ID string) (model.Currency, error)
	UpdateCurrency(ID string, params model.Currency) error
	DeleteCurrency(ID string) error
}

func NewCurrencyService(currency CurrencyRepo) *CurrencyService {
	return &CurrencyService{currency: currency}
}

func (cs *CurrencyService) CreateCurrency(params *schema.CreateCurrency) error {
	var data model.Currency

	data.Country = params.Country
	data.Currency = params.Currency
	data.CurrencyAbb = params.CurrencyAbb

	err := cs.currency.CreateCurrency(&data)
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

func (cs *CurrencyService) GetCurrencyByID(id string) (schema.GetCurrencyLists, error) {
	var response schema.GetCurrencyLists

	currency, err := cs.currency.GetCurrencyByID(id)
	if err != nil {
		return schema.GetCurrencyLists{}, errors.New("unable to get currency by id")
	}

	response.ID = currency.ID
	response.Country = currency.Country
	response.Currency = currency.Currency
	response.CurrencyAbb = currency.CurrencyAbb

	return response, nil
}

func (cs *CurrencyService) UpdateCurrency(ID string, params schema.GetCurrencyLists) error {
	var updateData model.Currency

	currency, err := cs.GetCurrencyByID(ID)
	if err != nil {
		log.Error("unable to update currency :%w ", err)
		return fmt.Errorf("unable to get currency by id")
	}

	updateData.Country = currency.Country
	if params.Country != "" {
		updateData.Country = params.Country
	}

	updateData.Currency = currency.Currency
	if params.Currency != "" {
		updateData.Currency = params.Currency
	}

	updateData.CurrencyAbb = currency.CurrencyAbb
	if params.Country != "" {
		updateData.CurrencyAbb = params.CurrencyAbb
	}

	err = cs.currency.UpdateCurrency(ID, updateData)
	if err != nil {
		return errors.New("unable to update currency")
	}
	return nil
}

func (cs *CurrencyService) DeleteCurrency(ID string) error {

	_, err := cs.currency.GetCurrencyByID(ID)
	if err != nil {
		return errors.New("unable to get currency by ID")
	}

	err = cs.currency.DeleteCurrency(ID)
	if err != nil {
		return fmt.Errorf("unable to delete this currency")
	}

	return nil
}
