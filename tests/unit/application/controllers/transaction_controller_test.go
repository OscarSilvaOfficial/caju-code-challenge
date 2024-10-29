package test

import (
	"caju-code-challenge/internal/application/controllers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecute_Cashin(t *testing.T) {
	mockService := &MockTransactionService{
		MakeCashinOperationFunc: func(accountId string, totalAmount float32, mcc string, merchant string) string {
			return "00"
		},
	}

	controller := controllers.NewTransactionController(mockService)
	input := controllers.TransactionInputData{
		AccountId:   "12345",
		TotalAmount: 100.0,
		Mcc:         "1234",
		Merchant:    "Merchant A",
		Cashin:      true,
	}

	output := controller.Execute(input)

	assert.Equal(t, "00", output.Code)
}

func TestExecute_Cashout(t *testing.T) {
	mockService := &MockTransactionService{
		MakeCashoutOperationFunc: func(accountId string, totalAmount float32, mcc string, merchant string) string {
			return "00"
		},
	}

	controller := controllers.NewTransactionController(mockService)
	input := controllers.TransactionInputData{
		AccountId:   "12345",
		TotalAmount: 100.0,
		Mcc:         "5678",
		Merchant:    "Merchant B",
		Cashin:      false,
	}

	output := controller.Execute(input)

	assert.Equal(t, "00", output.Code)
}
