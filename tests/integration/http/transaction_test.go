package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSubmitRoute(t *testing.T) {
	router := setupRouter()

	payload := []byte(`{
		"account": "123",
		"totalAmount": 100.00,
		"mcc": "5811",
		"merchant": "PADARIA DO ZE               SAO PAULO BR"
	}`)

	req, _ := http.NewRequest("POST", "/transactions", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	responseData, err := parseJsonResponse(w.Body.String())
	assert.NoError(t, err)
	assert.Equal(t, "123", responseData.Code.Account)
}
