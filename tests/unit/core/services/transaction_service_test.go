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

	assert.Equal(t, isAuthorized, true)
}

func TestDoCashoutOperationAndRefuseTransactoin(t *testing.T) {
	dbMock := &DbMock[services.TransactionService]{}

	service := services.NewTransactionService(dbMock)
	isAuthorized := service.MakeCashoutOperation(
		"account-id", 1100.00, "5412", "",
	)

	assert.Equal(t, isAuthorized, false)
}
