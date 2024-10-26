package test

import (
	"caju-code-challenge/internal/core/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoCashinOperation(t *testing.T) {
	wallet := entities.CreateWalletEntity("account-id", 1000.00, "FOOD")
	
	assert.Equal(t, wallet.NewCashInBalance(100.00), float32(1100.00))
	assert.Equal(t, wallet.GetCurrentBalance(), float32(1000.00))
}

func TestDoCashoutOperation(t *testing.T) {
	wallet := entities.CreateWalletEntity("account-id", 1000.00, "FOOD")
	
	assert.Equal(t, wallet.NewCashOutBalance(100.00), float32(900.00))
	assert.Equal(t, wallet.GetCurrentBalance(), float32(1000.00))
}
