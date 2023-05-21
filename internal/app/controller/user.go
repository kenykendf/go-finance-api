package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-finance/internal/app/schema"
	"github.com/go-finance/internal/pkg/handler"

	log "github.com/sirupsen/logrus"
)

type UserService interface {
	GetUsersLists(req *schema.UserPagination) ([]schema.GetUsersLists, error)
	CreateUser(req *schema.CreateUser) error
}

type UserController struct {
	userService UserService
}

func NewUserController(userService UserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) GetLists(ctx *gin.Context) {
	req := &schema.UserPagination{}
	req.Limit = ctx.GetInt("limit")
	req.Offset = ctx.GetInt("offset")

	data, err := uc.userService.GetUsersLists(req)
	if err != nil {
		log.Error(fmt.Errorf("unable to process GetUsersLists : %w", err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": data})
}

func (uc *UserController) Create(ctx *gin.Context) {
	req := &schema.CreateUser{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	err := uc.userService.CreateUser(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "user already created"})
}
