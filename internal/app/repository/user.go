package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/go-finance/internal/app/model"

	log "github.com/sirupsen/logrus"
)

type UserRepo struct {
	DB *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (ur *UserRepo) GetUsersLists(params model.UserPagination) ([]model.User, error) {
	var (
		limit        = params.Limit
		offset       = params.Offset
		users        []model.User
		sqlStatement = `
			SELECT id, username, email, fullname
			FROM users
			LIMIT $1
			OFFSET $2
		`
	)

	rows, err := ur.DB.Queryx(sqlStatement, limit, offset)
	if err != nil {
		log.Error(fmt.Errorf("error user repository - GetUsersLists : %w", err))
		return users, err
	}

	for rows.Next() {
		var user model.User
		err := rows.StructScan(&user)
		if err != nil {
			log.Error(fmt.Errorf("error user repository - GetUsersLists : %w", err))
		}
		users = append(users, user)
	}

	return users, nil
}

func (ur *UserRepo) GetUserByUsername(username string) (model.User, error) {
	var (
		user         model.User
		sqlStatement = `
			SELECT id, username, email, fullname
			FROM users
			WHERE username = $1
		`
	)

	err := ur.DB.QueryRowx(sqlStatement, username).StructScan(&user)
	if err != nil {
		log.Error(fmt.Errorf("error user repository - GetUserByUsername : %w", err))
		return user, err
	}

	return user, nil
}

func (ur *UserRepo) GetUserByEmail(email string) (model.User, error) {
	var (
		user         model.User
		sqlStatement = `
			SELECT id, username, email, fullname, password
			FROM users
			WHERE email = $1
		`
	)

	err := ur.DB.QueryRowx(sqlStatement, email).StructScan(&user)
	if err != nil {
		log.Error(fmt.Errorf("error user repository - GetUserByEmail : %w", err))
		return user, err
	}

	return user, nil
}

func (ur *UserRepo) GetUserByID(ID int) (model.User, error) {
	var (
		user         model.User
		sqlStatement = `
			SELECT id, username, email, fullname
			FROM users
			WHERE email = $1
		`
	)

	err := ur.DB.QueryRowx(sqlStatement, ID).StructScan(&user)
	if err != nil {
		log.Error(fmt.Errorf("error user repository - GetUserByID : %w", err))
		return user, err
	}

	return user, nil
}

func (ur *UserRepo) CreateUser(user model.User) error {
	var (
		sqlStatement = `
			INSERT INTO users (username,email,password,fullname)
			VALUES ($1,$2,$3,$4)
		`
	)

	_, err := ur.DB.Exec(sqlStatement, user.Username, user.Email, user.Password, user.Fullname)
	if err != nil {
		log.Error(fmt.Errorf("create user : %w", err))
		return err
	}

	return nil
}
