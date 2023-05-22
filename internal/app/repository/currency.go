package repository

import (
	"fmt"

	"github.com/go-finance/internal/app/model"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type CurrencyRepo struct {
	DB *sqlx.DB
}

func NewCurrencyRepo(db *sqlx.DB) *CurrencyRepo {
	return &CurrencyRepo{DB: db}
}

func (cr *CurrencyRepo) Create(params *model.Currency) error {
	var (
		sqlStatement = `
			INSERT INTO currency (country, currency, currency_abb)
			VALUES ($1, $2, $3)
		`
	)

	_, err := cr.DB.Exec(sqlStatement, params.Country, params.Currency, params.CurrencyAbb)
	if err != nil {
		log.Error(fmt.Errorf("error AuthRepo - Create : %w", err))
		return err
	}

	return nil
}

func (cr *CurrencyRepo) GetCurrenciesLists() ([]model.Currency, error) {
	var (
		sqlStatement = `
			SELECT id, country, currency, currency_abb
			FROM currency
		`
		currencies []model.Currency
	)

	rows, err := cr.DB.Queryx(sqlStatement)
	if err != nil {
		log.Error(fmt.Errorf("error currency repository - GetCurrenciesLists : %w", err))
		return currencies, err
	}

	for rows.Next() {
		var currency model.Currency
		err := rows.StructScan(&currency)
		if err != nil {
			log.Error(fmt.Errorf("error currency repository - GetCurrenciesLists : %w", err))
		}
		currencies = append(currencies, currency)
	}

	return currencies, nil
}
