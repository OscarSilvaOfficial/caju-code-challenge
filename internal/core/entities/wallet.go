package entities

import (
	"caju-code-challenge/internal/core/enums"
)

type Wallet struct {
	accountId string
	balance float32
	creditType enums.CreditType
}

func CreateWalletEntity(
	accountId string,
	balance float32,
	creditType enums.CreditType,
) Wallet {
	return Wallet{accountId, balance, creditType}
}

func (wallet *Wallet) GetCurrentBalance() float32 {
	return wallet.balance
}

func (credit *Wallet) NewCashInBalance(amount float32) float32 {
	return credit.balance + amount
}

func (credit *Wallet) NewCashOutBalance(amount float32) float32 {
	return credit.balance - amount
}