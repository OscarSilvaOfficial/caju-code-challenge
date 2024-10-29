package test

import (
	"caju-code-challenge/internal/core/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoCashoutOperationAndAcceptTransaction(t *testing.T) {
	dbMock := &DbMock[services.TransactionService]{}

	service := services.NewTransactionService(dbMock)
	isAuthorized := service.MakeCashoutOperation(
		"account-id", 900.00, "5412", "",
	)

	assert.Equal(t, isAuthorized, "00")
}

func TestDoCashoutOperationAndRefuseTransaction(t *testing.T) {
	dbMock := &DbMock[services.TransactionService]{}

	service := services.NewTransactionService(dbMock)
	isAuthorized := service.MakeCashoutOperation(
		"account-id", 2100.00, "5412", "",
	)

	assert.Equal(t, isAuthorized, "51")
}

func TestDoCashinOperationAndAcceptTransaction(t *testing.T) {
	dbMock := &DbMock[services.TransactionService]{}

	service := services.NewTransactionService(dbMock)
	isAuthorized := service.MakeCashinOperation(
		"account-id", 900.00, "5412", "",
	)

	assert.Equal(t, isAuthorized, "00")
}

func TestDoCashoutOperationAndRefuseTransactionByCategory(t *testing.T) {
	dbMock := &DbMock[services.TransactionService]{}

	service := services.NewTransactionService(dbMock)
	isAuthorized := service.MakeCashoutOperation(
		"account-id", 1001.00, "5811", "",
	)

	assert.Equal(t, isAuthorized, "51")
}
