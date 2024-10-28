package test

import (
	"caju-code-challenge/internal/infrastructure/adapters/db"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMongoDataCreation(t *testing.T) {
	dbName := "test"
	uri := os.Getenv("DB") 

	if uri == "" {
		uri = "mongodb://root:root@localhost:27017"
	}

	db, connectionError := db.NewMongoDB[Data](uri, dbName)

	assert.NoError(t, connectionError, "Connection error")

	insertData := Data{
		AccountId:   "account-1",
		TotalAmount: 1000.00,
		MCC:         "5411",
		Merchant:    "PADARIA DO ZE               SAO PAULO BR",
	}

	_, insertError := db.Insert(
		"transactions",
		insertData,
	)

	assert.NoError(t, insertError, "Insert should not return an error")

	transactions, findError := db.Find(
		"transactions", 
		map[string]interface{}{
			"accountid": "account-1",
		},
	)

	fmt.Println("Account ID: ", transactions[0].AccountId)

	assert.NoError(t, findError, "Find should not return an error")
	assert.Equal(t, transactions[0].AccountId, "account-1")
}
