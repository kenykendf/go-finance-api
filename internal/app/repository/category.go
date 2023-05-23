package repository

import (
	"fmt"

	"github.com/go-finance/internal/app/model"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type CategoryRepo struct {
	DB *sqlx.DB
}

func NewCategoryRepo(db *sqlx.DB) *CategoryRepo {
	return &CategoryRepo{DB: db}
}

func (cr *CategoryRepo) CreateCategory(params *model.Category) error {
	var (
		sqlStatement = `
			INSERT INTO category (name)
			VALUES ($1)
		`
	)

	_, err := cr.DB.Exec(sqlStatement, params.Name)
	if err != nil {
		log.Error(fmt.Errorf("error Category - CreateCategory : %w", err))
		return err
	}
	return nil
}

func (cr *CategoryRepo) GetCategoriesLists() ([]model.Category, error) {
	var (
		sqlStatement = `
			SELECT id, name
			FROM category
			WHERE deleted_at IS NULL
		`
		categories []model.Category
	)

	rows, err := cr.DB.Queryx(sqlStatement)
	if err != nil {
		log.Error(fmt.Errorf("error Category - GetCategoriesLists : %w", err))
		return categories, err
	}

	for rows.Next() {
		var category model.Category
		err := rows.StructScan(&category)
		if err != nil {
			log.Error(fmt.Errorf("error category repository - GetCurrenciesLists : %w", err))
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (cr *CategoryRepo) GetCategoryByID(id string) (model.Category, error) {
	var (
		sqlStatement = `
			SELECT id, name
			FROM category
			WHERE id = $1
			AND deleted_at IS NULL
		`
		category model.Category
	)

	err := cr.DB.QueryRowx(sqlStatement, id).StructScan(&category)
	if err != nil {
		log.Error(fmt.Errorf("error category repository - GetCategoryByID : %w", err))
		return category, err
	}

	return category, nil
}

func (cr *CategoryRepo) UpdateCategory(ID string, params model.Category) error {
	var (
		sqlStatement = `
			UPDATE category
			SET name = $2
			WHERE id = $1
			AND deleted_at IS NULL
		`
	)
	fmt.Println("ID = ", ID)
	fmt.Println("Params = ", params)
	_, err := cr.DB.Exec(sqlStatement, ID, params.Name)
	if err != nil {
		log.Error(fmt.Errorf("error category repository - UpdateCategory : %w", err))
		return err
	}

	return nil
}

func (cr *CategoryRepo) DeleteCategory(id string) error {
	var (
		sqlStatement = `
			UPDATE category
			SET deleted_at = now()
			WHERE id = $1
			AND deleted_at IS NULL
		`
	)

	_, err := cr.DB.Exec(sqlStatement, id)
	if err != nil {
		log.Error(fmt.Errorf("error category repository - DeleteCategory : %w", err))
		return err
	}

	return nil
}
