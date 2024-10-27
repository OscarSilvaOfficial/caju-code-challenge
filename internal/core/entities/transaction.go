package entities

type Transaction struct {
	accountId   string
	totalAmount float32
	mcc         string
	merchant    string
	cashin 			bool
}

func (transaction *Transaction) GetTotalAmount() float32 {
	return transaction.totalAmount
}

func (transaction *Transaction) IsCashin() bool {
	return transaction.cashin
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

func NewTransaction(
	accountId string,
	totalAmount float32,
	mcc string,
	merchant string,
	cashin bool,
) Transaction {
	return Transaction{accountId, totalAmount, mcc, merchant, cashin}
}
