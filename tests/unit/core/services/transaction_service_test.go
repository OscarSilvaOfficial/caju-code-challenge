package test

import (
	"caju-code-challenge/internal/core/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoCashoutOperationAndAcceptTransactoin(t *testing.T) {
	dbMock := &DbMock[services.TransactionService]{}

	service := services.NewTransactionService(dbMock)
	isAuthorized := service.MakeCashoutOperation(
		"account-id", 900.00, "5412", "",
	)

	assert.Equal(t, isAuthorized, "00")
}

func TestDoCashoutOperationAndRefuseTransactoin(t *testing.T) {
	dbMock := &DbMock[services.TransactionService]{}

	service := services.NewTransactionService(dbMock)
	isAuthorized := service.MakeCashoutOperation(
		"account-id", 2100.00, "5412", "",
	)

	assert.Equal(t, isAuthorized, "51")
}

func TestDoCashinOperationAndAcceptTransactoin(t *testing.T) {
	dbMock := &DbMock[services.TransactionService]{}

	service := services.NewTransactionService(dbMock)
	isAuthorized := service.MakeCashinOperation(
		"account-id", 900.00, "5412", "",
	)

	assert.Equal(t, isAuthorized, "00")
}
