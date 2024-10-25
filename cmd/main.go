package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()

	server.GET("/", func (ctx *gin.Context)  {
		ctx.JSON(200, gin.H{"Message": "Response"})
	})

	server.Run(":3000")
}