package main

import (
	routes "caju-code-challenge/internal/infrastructure/router"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	routes.Routes(server)
	server.Run(":8000")
}
