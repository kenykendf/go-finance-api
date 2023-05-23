package service

import (
	"errors"

	"github.com/go-finance/internal/app/model"
	"github.com/go-finance/internal/app/schema"
)

type CategoryService struct {
	categoryRepo CategoryRepo
}

type CategoryRepo interface {
	CreateCategory(params *model.Category) error
	GetCategoriesLists() ([]model.Category, error)
	GetCategoryByID(id string) (model.Category, error)
	UpdateCategory(id string, params model.Category) error
	DeleteCategory(id string) error
}

func NewCategoryService(categoryRepo CategoryRepo) *CategoryService {
	return &CategoryService{categoryRepo: categoryRepo}
}

func (cs *CategoryService) CreateCategory(params *schema.CreateCategory) error {
	var data model.Category

	data.Name = params.Name

	err := cs.categoryRepo.CreateCategory(&data)
	if err != nil {
		return errors.New("unable to create new category")
	}

	return nil
}

func (cs *CategoryService) GetCategoriesLists() ([]schema.GetCategories, error) {
	var response []schema.GetCategories

	categories, err := cs.categoryRepo.GetCategoriesLists()
	if err != nil {
		return response, errors.New("unable to get categories lists")
	}

	for _, v := range categories {
		var data schema.GetCategories
		data.ID = v.ID
		data.Name = v.Name

		response = append(response, data)
	}

	return response, nil
}

func (cs *CategoryService) GetCategoryByID(id string) (schema.GetCategories, error) {
	var response schema.GetCategories

	category, err := cs.categoryRepo.GetCategoryByID(id)
	if err != nil {
		return response, errors.New("unable to get categories by id")
	}

	response.ID = category.ID
	response.Name = category.Name

	return response, nil

}

func (cs *CategoryService) UpdateCategory(id string, params schema.CreateCategory) error {
	var updateData model.Category

	category, err := cs.categoryRepo.GetCategoryByID(id)
	if err != nil {
		return errors.New("unable to get categories by id")
	}

	updateData.Name = category.Name
	if params.Name != "" {
		updateData.Name = params.Name
	}

	err = cs.categoryRepo.UpdateCategory(id, updateData)
	if err != nil {
		return errors.New("unable to update category")
	}

	return nil
}

func (cs *CategoryService) DeleteCategory(id string) error {
	_, err := cs.categoryRepo.GetCategoryByID(id)
	if err != nil {
		return errors.New("unable to get category by id")
	}

	err = cs.categoryRepo.DeleteCategory(id)
	if err != nil {
		return errors.New("unable to get category by id")
	}

	return nil
}
