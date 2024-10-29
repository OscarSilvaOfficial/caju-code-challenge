package test

import (
	"caju-code-challenge/internal/infrastructure/web"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	server := gin.Default()
	web.Routes(server) 
	return server
}

func parseJsonResponse(jsonStr string) (*ApiResponse, error) {
	var response ApiResponse
	err := json.Unmarshal([]byte(jsonStr), &response)
	if err != nil {
			return nil, err
	}
	return &response, nil
}