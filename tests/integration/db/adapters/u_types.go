package test

type Data struct {
	AccountId   string  `json:"account_id"`
	TotalAmount float64 `json:"total_amount"`
	MCC         string  `json:"mcc"`
	Merchant    string  `json:"merchant"`
}