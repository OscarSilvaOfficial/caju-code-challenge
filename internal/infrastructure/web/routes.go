package web

import (
	"caju-code-challenge/internal/application/controllers"
	"caju-code-challenge/internal/infrastructure/factories"
	"caju-code-challenge/internal/infrastructure/web/parser"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(server *gin.Engine) {
	server.POST("/transactions", func(ctx *gin.Context) {
		var body parser.TransactionBody

		err := ctx.ShouldBindJSON(&body)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		transactionController := factories.NewTransactionController()

		data := transactionController.Execute(
			controllers.TransactionInputData{
				AccountId: body.Account,
				TotalAmount: float32(body.TotalAmount),
				Mcc: body.MCC,
				Merchant: body.Merchant,
				Cashin: body.Cashin,
			},
		)

		ctx.JSON(http.StatusOK, gin.H{"code": data.Code})
	})
}
