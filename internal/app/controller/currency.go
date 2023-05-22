package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-finance/internal/app/schema"
	"github.com/go-finance/internal/pkg/handler"

	log "github.com/sirupsen/logrus"
)

type CurrencyController struct {
	currency CurrencyService
}

type CurrencyService interface {
	CreateCurrency(params *schema.CreateCurrency) error
	GetCurrenciesLists() ([]schema.GetCurrencyLists, error)
	GetCurrencyByID(id string) (schema.GetCurrencyLists, error)
	UpdateCurrency(ID string, params schema.GetCurrencyLists) error
	DeleteCurrency(ID string) error
}

func NewCurrencyController(currency CurrencyService) *CurrencyController {
	return &CurrencyController{currency: currency}
}

func (cc *CurrencyController) CreateCurrency(ctx *gin.Context) {
	req := &schema.CreateCurrency{}

	if handler.BindAndCheck(ctx, req) {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, "Unable to Create Currency")
		return
	}

	err := cc.currency.CreateCurrency(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, "Unable to Create Currency")
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "Currency Created", nil)
}

func (cc *CurrencyController) GetCurrenciesLists(ctx *gin.Context) {
	data, err := cc.currency.GetCurrenciesLists()
	if err != nil {
		log.Error(fmt.Errorf("unable to process GetCurrenciesLists : %w", err))
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, "Unable To Get Currencies Lists")
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "Currency Created", data)
}

func (cc *CurrencyController) GetCurrencyByID(ctx *gin.Context) {
	currencyID := ctx.Param("id")

	data, err := cc.currency.GetCurrencyByID(currencyID)
	if err != nil {
		log.Error(fmt.Errorf("unable to process GetCurrencyByID : %w", err))
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success", data)
}

func (cc *CurrencyController) UpdateCurrency(ctx *gin.Context) {
	var params schema.GetCurrencyLists

	currencyID := ctx.Param("id")

	if handler.BindAndCheck(ctx, &params) {
		return
	}

	err := cc.currency.UpdateCurrency(currencyID, params)
	if err != nil {
		log.Error(fmt.Errorf("unable to process UpdateCurrency : %w", err))
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success", nil)
}

func (cc *CurrencyController) DeleteCurrency(ctx *gin.Context) {
	currencyID := ctx.Param("id")

	err := cc.currency.DeleteCurrency(currencyID)
	if err != nil {
		log.Error(fmt.Errorf("unable to process DeleteCurrency : %w", err))
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success", nil)
}
