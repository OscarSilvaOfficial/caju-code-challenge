package entities

type Transaction struct {
	accountId   string
	totalAmount float32
	mcc         string
	merchant    string
}

func CreateTransaction(
	accountId string,
	totalAmount float32,
	mcc string,
	merchant string,
) Transaction {
	return Transaction{
		accountId: accountId,
		totalAmount: totalAmount,
		mcc: mcc,
		merchant: merchant,
	}
}

func (transaction *Transaction) GetCreditType() string {
	if transaction.mcc == "5411" || transaction.mcc == "5412" {
		return "FOOD"
	}

	if transaction.mcc == "5811" || transaction.mcc == "5812" {
		return "MEAL"
	}

	return "CASH"
}

// func (transaction *Transaction) GetAccountId() string {
// 	return transaction.accountId
// }

// func (transaction *Transaction) GetTotalAmount() float32 {
// 	return transaction.totalAmount
// }

// func (transaction *Transaction) GetMCC() string {
// 	return transaction.mcc
// }

// func (transaction *Transaction) GetMerchant() string {
// 	return transaction.merchant
// }
