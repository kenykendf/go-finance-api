package schema

type GetTypeTransaction struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateTypeTransaction struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
