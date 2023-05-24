package repository

import (
	"fmt"

	"github.com/go-finance/internal/app/model"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type TypeTransactionRepo struct {
	DB *sqlx.DB
}

func NewTypeTransactionRepo(db *sqlx.DB) *TypeTransactionRepo {
	return &TypeTransactionRepo{DB: db}
}

func (ttr *TypeTransactionRepo) CreateTypeTransaction(params *model.TypeTransaction) error {
	var (
		sqlStatement = `
			INSERT INTO type_transaction (name, description)
			VALUES ($1, $2)
		`
	)

	_, err := ttr.DB.Exec(sqlStatement, params.Name, params.Description)
	if err != nil {
		log.Error(fmt.Errorf("error type_transaction repository - CreateTypeTransaction : %w", err))
		return err
	}

	return nil
}

func (ttr *TypeTransactionRepo) GetTypeTransactionsLists() ([]model.TypeTransaction, error) {
	var (
		sqlStatement = `
			SELECT id, name, description
			FROM type_transaction
			WHERE deleted_at IS NULL
		`
		typeTrans []model.TypeTransaction
	)

	rows, err := ttr.DB.Queryx(sqlStatement)
	if err != nil {
		log.Error(fmt.Errorf("error type_transaction repository - GetTypeTransactionsLists : %w", err))
		return typeTrans, err
	}

	for rows.Next() {
		var typeTran model.TypeTransaction
		err := rows.StructScan(&typeTran)
		if err != nil {
			log.Error(fmt.Errorf("error type_transaction repository - GetTypeTransactionsLists : %w", err))
		}
		typeTrans = append(typeTrans, typeTran)
	}

	return typeTrans, nil
}

func (ttr *TypeTransactionRepo) GetTypeTransactionByID(id string) (model.TypeTransaction, error) {
	var (
		sqlStatement = `
			SELECT id, name, description
			FROM type_transaction
			WHERE deleted_at IS NULL
			AND id = $1
		`
		typeTrans model.TypeTransaction
	)

	err := ttr.DB.QueryRowx(sqlStatement, id).StructScan(&typeTrans)
	if err != nil {
		log.Error(fmt.Errorf("error type_transaction repository - GetTypeTransByID : %w", err))
		return typeTrans, err
	}

	return typeTrans, nil
}

func (ttr *TypeTransactionRepo) UpdateTypeTransaction(ID string, params model.TypeTransaction) error {
	var (
		sqlStatement = `
			UPDATE type_transaction
			SET name = $2, description = $3, updated_at = now()
			WHERE id = $1
			AND deleted_at IS NULL
		`
	)

	_, err := ttr.DB.Exec(sqlStatement, ID, params.Name, params.Description)
	if err != nil {
		log.Error(fmt.Errorf("error type_transaction repository - UpdateTypeTransaction : %w", err))
		return err
	}

	return nil
}

func (ttr *TypeTransactionRepo) DeleteTypeTransacton(ID string) error {
	var (
		sqlStatement = `
			UPDATE type_transaction
			SET deleted_at = now()
			WHERE id = $1
			AND deleted_at IS NULL
		`
	)

	_, err := ttr.DB.Exec(sqlStatement, ID)
	if err != nil {
		log.Error(fmt.Errorf("error type_transaction repository - DeleteTypeTransacton : %w", err))
		return err
	}

	return nil
}
