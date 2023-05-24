package service

import (
	"errors"

	"github.com/go-finance/internal/app/model"
	"github.com/go-finance/internal/app/schema"
)

type TypeTransactionService struct {
	repo TypeTransactionRepo
}

type TypeTransactionRepo interface {
	CreateTypeTransaction(params *model.TypeTransaction) error
	GetTypeTransactionsLists() ([]model.TypeTransaction, error)
	GetTypeTransactionByID(id string) (model.TypeTransaction, error)
	UpdateTypeTransaction(id string, params model.TypeTransaction) error
	DeleteTypeTransacton(id string) error
}

func NewTypeTransactionService(repo TypeTransactionRepo) *TypeTransactionService {
	return &TypeTransactionService{repo: repo}

}

func (tts *TypeTransactionService) CreateTypeTransaction(params *schema.CreateTypeTransaction) error {
	var data model.TypeTransaction

	data.Name = params.Name
	data.Description = params.Description

	err := tts.repo.CreateTypeTransaction(&data)
	if err != nil {
		return errors.New("unable to create new type transaction")
	}

	return nil
}

func (tts *TypeTransactionService) GetTypeTransactionsLists() ([]schema.GetTypeTransaction, error) {
	var response []schema.GetTypeTransaction

	data, err := tts.repo.GetTypeTransactionsLists()
	if err != nil {
		return response, errors.New("unable to get type transactions lists")
	}

	for _, v := range data {
		var data schema.GetTypeTransaction
		data.ID = v.ID
		data.Name = v.Name
		data.Description = v.Description

		response = append(response, data)
	}

	return response, nil
}

func (tts *TypeTransactionService) GetTypeTransactionByID(id string) (schema.GetTypeTransaction, error) {
	var response schema.GetTypeTransaction

	data, err := tts.repo.GetTypeTransactionByID(id)
	if err != nil {
		return response, errors.New("unable to get type transactions by id")
	}

	response.ID = data.ID
	response.Name = data.Name
	response.Description = data.Description

	return response, nil

}

func (tts *TypeTransactionService) UpdateTypeTransaction(id string, params schema.CreateTypeTransaction) error {
	var updateData model.TypeTransaction

	data, err := tts.repo.GetTypeTransactionByID(id)
	if err != nil {
		return errors.New("unable to get type transactions by id")
	}

	updateData.Name = data.Name
	if params.Name != "" {
		updateData.Name = params.Name
	}
	updateData.Description = data.Description
	if params.Description != "" {
		updateData.Description = params.Description
	}

	err = tts.repo.UpdateTypeTransaction(id, updateData)
	if err != nil {
		return errors.New("unable to update type transaction")
	}

	return nil
}

func (tts *TypeTransactionService) DeleteTypeTransacton(id string) error {
	_, err := tts.repo.GetTypeTransactionByID(id)
	if err != nil {
		return errors.New("unable to get transaction by id")
	}

	err = tts.repo.DeleteTypeTransacton(id)
	if err != nil {
		return errors.New("unable to get transaction by id")
	}

	return nil
}
