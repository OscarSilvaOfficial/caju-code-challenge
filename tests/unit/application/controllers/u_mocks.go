package test

type MockTransactionService struct {
	MakeCashinOperationFunc func(accountId string, totalAmount float32, mcc string, merchant string) string
	MakeCashoutOperationFunc func(accountId string, totalAmount float32, mcc string, merchant string) string
}

func (m *MockTransactionService) MakeCashinOperation(accountId string, totalAmount float32, mcc string, merchant string) string {
	return m.MakeCashinOperationFunc(accountId, totalAmount, mcc, merchant)
}

func (m *MockTransactionService) MakeCashoutOperation(accountId string, totalAmount float32, mcc string, merchant string) string {
	return m.MakeCashoutOperationFunc(accountId, totalAmount, mcc, merchant)
}