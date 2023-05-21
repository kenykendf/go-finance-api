package schema

type UserPagination struct {
	Limit  int
	Offset int
}

type GetUsersLists struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	// created_at string `json:"created_at"`
}

type CreateUser struct {
	Username string `json:"username" validate:"required,max=20"`
	Email    string `json:"email" validate:"required,email"`
	Fullname string `json:"fullname" validate:"required,max=100"`
	Password string `json:"password" validate:"required,min=8,alphanum"`
}
