package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteCashoutOperationWithOkStatus(t *testing.T) {
	router := setupRouter()

	payload := []byte(`{
		"account": "integration-test-1",
		"totalAmount": 100.00,
		"mcc": "5811",
		"merchant": "PADARIA DO ZE               SAO PAULO BR",
		"cashin": true
	}`)

	cashinReq, _ := http.NewRequest("POST", "/transactions", bytes.NewBuffer(payload))
	cashinReq.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, cashinReq)

	assert.Equal(t, http.StatusOK, w.Code)

	cashinResponseBody := w.Body.String()
	cashinResponseData, err := parseJsonResponse(cashinResponseBody)
	assert.NoError(t, err)
	assert.Equal(t, "00", cashinResponseData.Code)

	payload = []byte(`{
		"account": "integration-test-1",
		"totalAmount": 100.00,
		"mcc": "5811",
		"merchant": "PADARIA DO ZE               SAO PAULO BR",
		"cashin": false
	}`)

	cashoutReq, _ := http.NewRequest("POST", "/transactions", bytes.NewBuffer(payload))
	cashoutReq.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	router.ServeHTTP(w, cashoutReq)

	cashoutResponseData, err := parseJsonResponse(w.Body.String())
	assert.NoError(t, err)
	assert.Equal(t, "00", cashoutResponseData.Code)
}

func TestExecuteCashoutOperationWithBlockStatus(t *testing.T) {
	router := setupRouter()

	payload := []byte(`{
		"account": "integration-test-2",
		"totalAmount": 100.00,
		"mcc": "5811",
		"merchant": "PADARIA DO ZE               SAO PAULO BR",
		"cashin": true
	}`)

	cashinReq, _ := http.NewRequest("POST", "/transactions", bytes.NewBuffer(payload))
	cashinReq.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, cashinReq)

	assert.Equal(t, http.StatusOK, w.Code)

	cashinResponseBody := w.Body.String()
	cashinResponseData, err := parseJsonResponse(cashinResponseBody)
	assert.NoError(t, err)
	assert.Equal(t, "00", cashinResponseData.Code)

	payload = []byte(`{
		"account": "integration-test-2",
		"totalAmount": 1000.00,
		"mcc": "5811",
		"merchant": "PADARIA DO ZE               SAO PAULO BR",
		"cashin": false
	}`)

	cashoutReq, _ := http.NewRequest("POST", "/transactions", bytes.NewBuffer(payload))
	cashoutReq.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	router.ServeHTTP(w, cashoutReq)

	cashoutResponseData, err := parseJsonResponse(w.Body.String())
	assert.NoError(t, err)
	assert.Equal(t, "51", cashoutResponseData.Code)
}
