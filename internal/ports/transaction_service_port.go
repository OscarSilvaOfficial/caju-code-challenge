package ports

type TransactionServicePort interface {
	MakeCashoutOperation(
		accountId string,
		totalAmount float32,
		mcc string,
		merchant string,
	) string

	MakeCashinOperation(
		accountId string,
		totalAmount float32,
		mcc string,
		merchant string,
	) string
}