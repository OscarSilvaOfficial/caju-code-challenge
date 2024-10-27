package services

import (
	"caju-code-challenge/internal/core/entities"
	"caju-code-challenge/internal/ports/output"
	"caju-code-challenge/internal/utils"
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

	transactions, _ := transactionService.findUserTransactions(accountId)

	calculatedValue := transactionService.calculateDebits(
		append(transactions, cashoutTransaction),
		cashoutTransaction.GetCreditType(),
	)

	return transactionService.isAuthorized(calculatedValue)
}

func (transactionService *TransactionService) isAuthorized(balance float32) bool {
	isAuthorized := balance >= 0
	return isAuthorized
}

func (transactionService *TransactionService) findUserTransactions(accountId string) ([]entities.Transaction, error) {
	userTransactions, err := transactionService.db.Find(
		"transactions",
		map[string]interface{}{
			"accountId": accountId,
		},
	)

	if err != nil {
		genericTransaction := entities.NewTransaction("", 0.00, "", "", true)
		return []entities.Transaction{genericTransaction}, err
	}

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

	return transactions, nil
}

func (transactionService *TransactionService) calculateDebits(transactions []entities.Transaction, transactionType string) float32 {
	return utils.Reduce(transactions, 0.0, func(accumulator float32, value entities.Transaction) float32 {
		isSameCreditType := value.GetCreditType() == transactionType

		if isSameCreditType && value.IsCashin() {
			return accumulator + value.GetTotalAmount()
		}

		if isSameCreditType && !value.IsCashin() {
			return accumulator - value.GetTotalAmount()
		}
		
		return accumulator
	})
}

func NewTransactionService(db output.DatabasePort[TransactionOutputData]) *TransactionService {
	return &TransactionService{db}
}
