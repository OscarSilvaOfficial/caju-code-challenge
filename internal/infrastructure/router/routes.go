package routes

import (
	"caju-code-challenge/internal/infrastructure/router/parser"
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

		ctx.JSON(http.StatusOK, gin.H{
			"code": body,
		})
	})
}
