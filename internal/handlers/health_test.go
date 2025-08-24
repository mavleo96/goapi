package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	HealthCheck(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("status code = %v, want %v", w.Code, http.StatusOK)
	}

	var response HealthResponse
	json.Unmarshal(w.Body.Bytes(), &response)
	if response.Status != "healthy" {
		t.Errorf("status = %v, want healthy", response.Status)
	}
}
