package main

import (
	routes "caju-code-challenge/internal/infrastructure/router"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	server := gin.Default()
	routes.Routes(server) 
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