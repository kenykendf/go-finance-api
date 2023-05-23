package schema

type GetCategories struct {
	ID   string `db:"id"`
	Name string `db:"name"`
}

type CreateCategory struct {
	Name string `json:"name"`
}
