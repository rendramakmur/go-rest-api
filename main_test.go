package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Home(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()

	Server().ServeHTTP(response, request)
	expectedResponse := `{"message":"Hello, welcome","status":200}`
	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		t.Log(err)
	}

	assert.Equal(t, 200, response.Code, "Invalid response code")
	assert.Equal(t, expectedResponse, string(bytes.TrimSpace(responseBody)))
	t.Log(response.Body.String())
}
