package middleware

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mavleo96/goapi/api"
)

func mockHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))
}

func TestAuthorizationSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/account/coins?username=alex", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "123ABC")

	w := httptest.NewRecorder()
	handler := Authorization(http.HandlerFunc(mockHandler))
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("status code = %v, want %v", w.Code, http.StatusOK)
	}
}

func TestAuthorizationFailure(t *testing.T) {
	req, err := http.NewRequest("GET", "/account/coins?username=alex", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "invalid")

	w := httptest.NewRecorder()
	handler := Authorization(http.HandlerFunc(mockHandler))
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("status code = %v, want %v", w.Code, http.StatusBadRequest)
	}

	var response api.Error
	json.Unmarshal(w.Body.Bytes(), &response)
	if response.Message != ErrUnauthorized.Error() {
		t.Errorf("error message = %v, want %v", response.Message, ErrUnauthorized.Error())
	}
}
