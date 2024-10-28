package main

type ResponseData struct {
	Account     string  `json:"account"`
	TotalAmount float64 `json:"totalAmount"`
	MCC         string  `json:"mcc"`
	Merchant    string  `json:"merchant"`
}

type ApiResponse struct {
	Code    ResponseData `json:"code"`
}