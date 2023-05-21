package service

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/go-finance/internal/app/model"
	"github.com/go-finance/internal/app/schema"
)

type UserRepo interface {
	GetUsersLists(users model.UserPagination) ([]model.User, error)
	GetUserByUsername(username string) (model.User, error)
	GetUserByEmail(email string) (model.User, error)
	GetUserByID(ID int) (model.User, error)
	CreateUser(user model.User) error
}

type UserService struct {
	userRepo UserRepo
}

func NewUserService(userRepo UserRepo) *UserService {
	return &UserService{userRepo: userRepo}
}

func (us *UserService) GetUsersLists(req *schema.UserPagination) ([]schema.GetUsersLists, error) {
	var response []schema.GetUsersLists

	dbSearch := model.UserPagination{}
	if req.Limit == 0 {
		req.Limit = 10
	}
	if req.Offset == 0 {
		req.Offset = 0
	}
	dbSearch.Limit = req.Limit
	dbSearch.Offset = req.Offset

	users, err := us.userRepo.GetUsersLists(dbSearch)
	if err != nil {
		return nil, errors.New("unable to get users lists")
	}

	for _, value := range users {
		var data schema.GetUsersLists
		data.ID = value.ID
		data.Fullname = value.Fullname
		data.Username = value.Username
		data.Email = value.Email

		response = append(response, data)
	}

	return response, nil
}

func (us *UserService) CreateUser(req *schema.CreateUser) error {
	var insertData model.User

	insertData.Email = req.Email
	insertData.Fullname = req.Fullname
	insertData.Username = req.Username

	pass, _ := us.hashPassword(req.Password)
	insertData.Password = pass

	err := us.userRepo.CreateUser(insertData)
	if err != nil {
		return errors.New("unable to create user")
	}

	return nil
}

func (us *UserService) hashPassword(password string) (string, error) {
	bytePass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytePass), nil

}
