package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-finance/internal/app/schema"
	"github.com/go-finance/internal/pkg/handler"
	log "github.com/sirupsen/logrus"
)

type TypeTransactionController struct {
	typeTransactionService TypeTransactionService
}

type TypeTransactionService interface {
	CreateTypeTransaction(params *schema.CreateTypeTransaction) error
	GetTypeTransactionsLists() ([]schema.GetTypeTransaction, error)
	GetTypeTransactionByID(id string) (schema.GetTypeTransaction, error)
	UpdateTypeTransaction(id string, params schema.CreateTypeTransaction) error
	DeleteTypeTransacton(id string) error
}

func NewTypeTransactionController(typeService TypeTransactionService) *TypeTransactionController {
	return &TypeTransactionController{typeTransactionService: typeService}
}

func (ttc *TypeTransactionController) CreateTypeTransaction(ctx *gin.Context) {
	req := &schema.CreateTypeTransaction{}

	if handler.BindAndCheck(ctx, req) {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, "unable to create type transaction")
		return
	}

	err := ttc.typeTransactionService.CreateTypeTransaction(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, "unable to create type transaction")
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "New Type Transaction Created", nil)
}

func (ttc *TypeTransactionController) GetTypeTransactionsLists(ctx *gin.Context) {
	data, err := ttc.typeTransactionService.GetTypeTransactionsLists()
	if err != nil {
		log.Error(fmt.Errorf("unable to process GetTypeTransactionsLists : %w", err))
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, "Unable To Get Type Transactions Lists")
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success", data)
}

func (ttc *TypeTransactionController) GetTypeTransactionByID(ctx *gin.Context) {
	categoryID := ctx.Param("id")

	data, err := ttc.typeTransactionService.GetTypeTransactionByID(categoryID)
	if err != nil {
		log.Error(fmt.Errorf("unable to process GetTypeTransactionByID : %w", err))
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, "Unable To Get Type Transaction By ID")
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success", data)
}

func (ttc *TypeTransactionController) UpdateTypeTransaction(ctx *gin.Context) {
	categoryID := ctx.Param("id")
	req := schema.CreateTypeTransaction{}

	if handler.BindAndCheck(ctx, &req) {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, "Unable To Update Type Transaction")
		return
	}

	err := ttc.typeTransactionService.UpdateTypeTransaction(categoryID, req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, "Unable To Update Type Transaction")
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "Category Updated", nil)
}

func (ttc *TypeTransactionController) DeleteTypeTransacton(ctx *gin.Context) {
	categoryID := ctx.Param("id")

	err := ttc.typeTransactionService.DeleteTypeTransacton(categoryID)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, "Unable To Delete Type Transaction")
	}
}
