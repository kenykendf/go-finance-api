package model

type Currency struct {
	ID          string `db:"id"`
	Country     string `db:"country"`
	Currency    string `db:"currency"`
	CurrencyAbb string `db:"currency_abb"`
}
