package services

import (
	"caju-code-challenge/internal/core/entities"
	"caju-code-challenge/internal/ports"
	"caju-code-challenge/internal/utils"
)

type TransactionOutputData struct {
	AccountId   string  `json:"accountid"`
	TotalAmount float32 `json:"totalamount"`
	Mcc         string  `json:"mcc"`
	Merchant    string  `json:"merchant"`
	Cashin      bool    `json:"cashin"`
}

type TransactionService struct {
	db ports.DatabasePort[TransactionOutputData]
}

func (transactionService *TransactionService) MakeCashoutOperation(
	accountId string,
	totalAmount float32,
	mcc string,
	merchant string,
) string {
	cashoutTransaction := entities.NewTransaction(
		accountId,
		totalAmount,
		mcc,
		merchant,
		false,
	)

	transactions, err := transactionService.findUserTransactions(accountId)

	if err != nil || len(transactions) == 0 {
		return "07"
	}

	calculatedValue := transactionService.calculateDebits(
		append(transactions, cashoutTransaction),
		cashoutTransaction.GetCreditType(),
	)

	if transactionService.isAuthorized(calculatedValue) {
		transactionService.db.Insert("transactions", TransactionOutputData{
			accountId,
			totalAmount,
			mcc,
			merchant,
			false,
		})

		return "00"
	}

	return "51"
}

func (transactionService *TransactionService) MakeCashinOperation(
	accountId string,
	totalAmount float32,
	mcc string,
	merchant string,
) string {
	_, err := transactionService.db.Insert("transactions", TransactionOutputData{
		accountId,
		totalAmount,
		mcc,
		merchant,
		true,
	})

	if err != nil {
		return "07"
	}

	return "00"
}

func (transactionService *TransactionService) isAuthorized(balance float32) bool {
	isAuthorized := balance >= 0
	return isAuthorized
}

func (transactionService *TransactionService) findUserTransactions(accountId string) ([]entities.Transaction, error) {
	userTransactions, err := transactionService.db.Find(
		"transactions",
		map[string]interface{}{
			"accountid": accountId,
		},
	)

	if err != nil {
		return nil, err
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

func NewTransactionService(db ports.DatabasePort[TransactionOutputData]) *TransactionService {
	return &TransactionService{db}
}
