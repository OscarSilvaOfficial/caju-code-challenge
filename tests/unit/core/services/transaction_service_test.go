package test

import (
	"caju-code-challenge/internal/core/services"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type DbMock[Data any] struct{}

func (db *DbMock[Data]) Find(collectionOrTable string, where map[string]interface{}) ([]services.TransactionOutputData, error) {
	data := `[
		{
			"id": "id",
			"accountId": "account-id",
			"totalAmount": 1000.00,
			"mcc": "5411",
			"merchant": "TEST",
			"cashin": true
		}
	]`

	var transactions []services.TransactionOutputData

	json.Unmarshal([]byte(data), &transactions)

	return transactions, nil
}

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
