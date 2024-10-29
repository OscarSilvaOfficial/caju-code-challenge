package test

import (
	"caju-code-challenge/internal/core/services"
	"encoding/json"
)

type DbMock[Data any] struct{}

func (db *DbMock[Data]) Find(collectionOrTable string, where map[string]interface{}) ([]services.TransactionOutputData, error) {
	data := `[
		{
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

func (db *DbMock[Data]) Insert(collectionOrTable string, data services.TransactionOutputData) (interface{}, error) {
	var transaction services.TransactionOutputData
	return transaction, nil
}
