package factories

import (
	"caju-code-challenge/internal/application/controllers"
	"caju-code-challenge/internal/core/services"
	"caju-code-challenge/internal/infrastructure/adapters/db"
	"fmt"
	"os"
)

func NewTransactionController() *controllers.TransactionController {
	dbName := os.Getenv("DB_NAME")
	connectionString := os.Getenv("DB_CONNECTION")
	database, err := db.NewMongoDB[services.TransactionOutputData](connectionString, dbName)

	if err != nil {
		fmt.Errorf("Mongo connection error", err)
		return nil
	}

	transactionService := services.NewTransactionService(database)

	return controllers.NewTransactionController(transactionService)
}