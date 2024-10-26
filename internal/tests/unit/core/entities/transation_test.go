package test

import (
	"caju-code-challenge/internal/core/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCreditTypeForFoodMcc(t *testing.T) {
	transaction1 := entities.CreateTransaction(
		"account-1",
		661.00,
		"5411",
		"BARRACA DO LÚCIO M               RS BR",
	)
	
	transaction2 := entities.CreateTransaction(
		"account-1",
		400.00,
		"5412",
		"MERCADO DO HERMANOTEU               GODA",
	)
	
	assert.Equal(t, transaction1.GetCreditType(), "FOOD")
	assert.Equal(t, transaction2.GetCreditType(), "FOOD")
}

func TestGetCreditTypeForMealMcc(t *testing.T) {
	transaction1 := entities.CreateTransaction(
		"account-1",
		661.00,
		"5811",
		"BARRACA DO LÚCIO M               RS BR",
	)

	transaction2 := entities.CreateTransaction(
		"account-1",
		400.00,
		"5812",
		"MERCADO DO HERMANOTEU               GODA",
	)

	assert.Equal(t, transaction1.GetCreditType(), "MEAL")
	assert.Equal(t, transaction2.GetCreditType(), "MEAL")
}

func TestGetCreditTypeForCashMcc(t *testing.T) {
	transaction := entities.CreateTransaction(
		"account-1",
		661.00,
		"1111",
		"BARRACA DO LÚCIO M               RS BR",
	)

	assert.Equal(t, transaction.GetCreditType(), "CASH")
}
