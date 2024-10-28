package parser

type TransactionBody struct {
	Account     string  `json:"account" binding:"required"`
	TotalAmount float64 `json:"totalAmount" binding:"required"`
	MCC         string  `json:"mcc" binding:"required"`
	Merchant    string  `json:"merchant" binding:"required"`
}