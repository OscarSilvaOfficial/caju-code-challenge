package parser

type TransactionBody struct {
	Account     string  `json:"account" binding:"required"`
	TotalAmount float32 `json:"totalAmount" binding:"required"`
	MCC         string  `json:"mcc" binding:"required"`
	Merchant    string  `json:"merchant" binding:"required"`
	Cashin   	 	bool  `json:"cashin"`
}