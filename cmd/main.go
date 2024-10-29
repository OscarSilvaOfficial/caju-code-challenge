package main

import (
	"caju-code-challenge/internal/infrastructure/web"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	web.Routes(server)
	server.Run(":8000")
}
