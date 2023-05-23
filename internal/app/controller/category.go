package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-finance/internal/app/schema"
	"github.com/go-finance/internal/pkg/handler"
	log "github.com/sirupsen/logrus"
)

type CategoryController struct {
	categoryService CategoryService
}

type CategoryService interface {
	CreateCategory(params *schema.CreateCategory) error
	GetCategoriesLists() ([]schema.GetCategories, error)
	GetCategoryByID(id string) (schema.GetCategories, error)
	UpdateCategory(id string, params schema.CreateCategory) error
	DeleteCategory(id string) error
}

func NewCategoryController(categoryService CategoryService) *CategoryController {
	return &CategoryController{categoryService: categoryService}
}

func (cc *CategoryController) CreateCategory(ctx *gin.Context) {
	req := &schema.CreateCategory{}

	if handler.BindAndCheck(ctx, req) {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, "unable to create category")
		return
	}

	err := cc.categoryService.CreateCategory(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, "unable to create category")
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "Category Created", nil)
}

func (cc *CategoryController) GetCategoriesLists(ctx *gin.Context) {
	data, err := cc.categoryService.GetCategoriesLists()
	if err != nil {
		log.Error(fmt.Errorf("unable to process GetCategoriesLists : %w", err))
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, "Unable To Get Categories Lists")
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success", data)
}

func (cc *CategoryController) GetCategoryByID(ctx *gin.Context) {
	categoryID := ctx.Param("id")

	data, err := cc.categoryService.GetCategoryByID(categoryID)
	if err != nil {
		log.Error(fmt.Errorf("unable to process GetCategoryByID : %w", err))
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, "Unable To Get Category By ID")
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success", data)
}

func (cc *CategoryController) UpdateCategory(ctx *gin.Context) {
	categoryID := ctx.Param("id")
	req := schema.CreateCategory{}

	if handler.BindAndCheck(ctx, &req) {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, "Unable To Update Category")
		return
	}

	err := cc.categoryService.UpdateCategory(categoryID, req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, "Unable To Update Category")
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "Category Updated", nil)
}

func (cc *CategoryController) DeleteCategory(ctx *gin.Context) {
	categoryID := ctx.Param("id")

	err := cc.categoryService.DeleteCategory(categoryID)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, "Unable To Delete Category")
	}
}
