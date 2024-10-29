package controllers

import (
	"caju-code-challenge/internal/ports"
)

type TransactionInputData struct {
	AccountId   string 
	TotalAmount float32
	Mcc         string 
	Merchant    string 
	Cashin      bool   
}

type ExecuteOutput struct {
	Code	string	`json:"code"`
}

type TransactionController struct {
	transactionService ports.TransactionServicePort
}

func (transactionController *TransactionController) Execute(data TransactionInputData) ExecuteOutput {
	if data.Cashin {
		code := transactionController.transactionService.MakeCashinOperation(data.AccountId, data.TotalAmount, data.Mcc, data.Merchant)
		return ExecuteOutput{code}
	}

	code := transactionController.transactionService.MakeCashoutOperation(data.AccountId, data.TotalAmount, data.Mcc, data.Merchant)
	return ExecuteOutput{code}
}

func NewTransactionController(transactionService ports.TransactionServicePort) *TransactionController {
	return &TransactionController{transactionService}
}