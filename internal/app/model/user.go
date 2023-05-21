package model

type UserPagination struct {
	Limit  int
	Offset int
}

type User struct {
	ID       int    `db:"id"`
	Username string `db:"username"`
	Email    string `db:"email"`
	Fullname string `db:"fullname"`
	Password string `db:"password"`
}
