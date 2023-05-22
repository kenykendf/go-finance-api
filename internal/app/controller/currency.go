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
	Create(params *schema.CreateCurrency) error
	GetCurrenciesLists() ([]schema.GetCurrencyLists, error)
}

func NewCurrencyController(currency CurrencyService) *CurrencyController {
	return &CurrencyController{currency: currency}
}

func (cc *CurrencyController) Create(ctx *gin.Context) {
	req := &schema.CreateCurrency{}

	if handler.BindAndCheck(ctx, req) {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, "Unable to Create Currency")
		return
	}

	err := cc.currency.Create(req)
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
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "Currency Created", data)
}
