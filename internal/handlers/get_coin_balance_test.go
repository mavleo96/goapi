package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mavleo96/goapi/api"
)

func TestGetCoinBalanceSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/account/coins?username=alex", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "123ABC")

	w := httptest.NewRecorder()
	GetCoinBalance(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("status code = %v, want %v", w.Code, http.StatusOK)
	}

	var response api.CoinBalanceResponse
	json.Unmarshal(w.Body.Bytes(), &response)
	if response.Balance != 100 {
		t.Errorf("balance = %v, want 100", response.Balance)
	}
}

func TestGetCoinBalanceMissingUsername(t *testing.T) {
	req, err := http.NewRequest("GET", "/account/coins", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "123ABC")

	w := httptest.NewRecorder()
	GetCoinBalance(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("status code = %v, want %v", w.Code, http.StatusInternalServerError)
	}
}
