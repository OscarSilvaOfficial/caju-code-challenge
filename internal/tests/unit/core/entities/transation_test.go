package test

import (
	"caju-code-challenge/internal/core/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCreditTypeForFoodMcc(t *testing.T) {
	transaction1 := entities.CreateTransactionEntity(
		"account-1",
		661.00,
		"5411",
		"BARRACA DO LÚCIO M               RS BR",
		false,
	)
	
	transaction2 := entities.CreateTransactionEntity(
		"account-1",
		400.00,
		"5412",
		"MERCADO DO HERMANOTEU               GODA",
		false,
	)
	
	assert.Equal(t, transaction1.GetCreditType(), "FOOD")
	assert.Equal(t, transaction2.GetCreditType(), "FOOD")
}

func TestGetCreditTypeForMealMcc(t *testing.T) {
	transaction1 := entities.CreateTransactionEntity(
		"account-1",
		661.00,
		"5811",
		"BARRACA DO LÚCIO M               RS BR",
		false,
	)

	transaction2 := entities.CreateTransactionEntity(
		"account-1",
		400.00,
		"5812",
		"MERCADO DO HERMANOTEU               GODA",
		false,
	)

	assert.Equal(t, transaction1.GetCreditType(), "MEAL")
	assert.Equal(t, transaction2.GetCreditType(), "MEAL")
}

func TestGetCreditTypeForCashMcc(t *testing.T) {
	transaction := entities.CreateTransactionEntity(
		"account-1",
		661.00,
		"1111",
		"BARRACA DO LÚCIO M               RS BR",
		false,
	)

	assert.Equal(t, transaction.GetCreditType(), "CASH")
}
