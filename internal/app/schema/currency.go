package schema

type GetCurrencyLists struct {
	ID          string `json:"id"`
	Country     string `json:"country"`
	Currency    string `json:"currency"`
	CurrencyAbb string `json:"currency_abb"`
}

type CreateCurrency struct {
	Country     string `json:"country"`
	Currency    string `json:"currency"`
	CurrencyAbb string `json:"currency_abb"`
}
