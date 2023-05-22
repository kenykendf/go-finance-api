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

func (cr *CurrencyRepo) CreateCurrency(params *model.Currency) error {
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
			WHERE deleted_at IS NULL
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

func (cr *CurrencyRepo) GetCurrencyByID(id string) (model.Currency, error) {
	var (
		sqlStatement = `
			SELECT id, country, currency, currency_abb
			FROM currency
			WHERE id = $1
			AND deleted_at IS NULL
		`
		currency model.Currency
	)

	err := cr.DB.QueryRowx(sqlStatement, id).StructScan(&currency)
	if err != nil {
		log.Error(fmt.Errorf("error currency repository - GetCurrenciesLists : %w", err))
		return currency, err
	}

	return currency, nil
}

func (cr *CurrencyRepo) UpdateCurrency(ID string, params model.Currency) error {
	var (
		sqlStatement = `
			UPDATE currency
			SET country = $2, currency = $3, currency_abb = $4, updated_at = now()
			WHERE id = $1
			AND deleted_at IS NULL
		`
	)

	_, err := cr.DB.Exec(sqlStatement, ID, params.Country, params.Currency, params.CurrencyAbb)
	if err != nil {
		log.Error(fmt.Errorf("error currency repository - UpdateCurrency : %w", err))
		return err
	}

	return nil
}

func (cr *CurrencyRepo) DeleteCurrency(ID string) error {
	var (
		sqlStatement = `
			UPDATE currency
			SET deleted_at = now()
			WHERE id = $1
			AND deleted_at IS NULL
		`
	)

	_, err := cr.DB.Exec(sqlStatement, ID)
	if err != nil {
		log.Error(fmt.Errorf("error currency repository - DeleteCurrency : %w", err))
		return err
	}

	return nil
}
