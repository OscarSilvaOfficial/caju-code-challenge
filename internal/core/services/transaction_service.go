package services

import (
	"caju-code-challenge/internal/core/entities"
	"caju-code-challenge/internal/ports/output"
)

type TransactionOutputData struct {
	Id          string  `json:"id"`
	AccountId   string  `json:"accountId"`
	TotalAmount float32 `json:"totalAmount"`
	Mcc         string  `json:"mcc"`
	Merchant    string  `json:"merchant"`
	Cashin      bool    `json:"cashin"`
}

type TransactionService struct {
	db output.DatabasePort[TransactionOutputData]
}

func (transactionService *TransactionService) MakeCashoutOperation(
	accountId string,
	totalAmount float32,
	mcc string,
	merchant string,
) bool {
	cashoutTransaction := entities.NewTransaction(
		accountId,
		totalAmount,
		mcc,
		merchant,
		false,
	)

	userTransactions, _ := transactionService.db.Find(
		"transactions",
		map[string]interface{}{
			"accountId": accountId,
		},
	)

	var transactions []entities.Transaction

	for _, value := range userTransactions {
		entity := entities.NewTransaction(
			value.AccountId, value.TotalAmount, value.Mcc, value.Merchant, value.Cashin,
		)

		transactions = append(
			transactions,
			entity,
		)
	}

	finalValue := transactionService.calculateDebits(
		append(transactions, cashoutTransaction),
		cashoutTransaction.GetCreditType(),
	)

	isAuthorized := finalValue >= 0

	return isAuthorized
}

func (transactionService *TransactionService) calculateDebits(transactions []entities.Transaction, transactionType string) float32 {
	finalValue := float32(0.00)

	for _, value := range transactions {
		isEqualTransactionType := value.GetCreditType() == transactionType

		if isEqualTransactionType && value.IsCashin() {
			finalValue += value.GetTotalAmount()
		}

		if isEqualTransactionType && !value.IsCashin() {
			finalValue -= value.GetTotalAmount()
		}
	}

	return finalValue
}

func NewTransactionService(db output.DatabasePort[TransactionOutputData]) *TransactionService {
	return &TransactionService{db}
}
