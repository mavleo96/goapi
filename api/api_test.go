package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriteError(t *testing.T) {
	w := httptest.NewRecorder()
	writeError(w, "test error", http.StatusBadRequest)

	if w.Code != http.StatusBadRequest {
		t.Errorf("status code = %v, want %v", w.Code, http.StatusBadRequest)
	}

	var response Error
	json.Unmarshal(w.Body.Bytes(), &response)
	if response.Message != "test error" {
		t.Errorf("message = %v, want test error", response.Message)
	}
}

func TestRequestErrorHandler(t *testing.T) {
	w := httptest.NewRecorder()
	testError := errors.New("test error")
	RequestErrorHandler(w, testError)

	if w.Code != http.StatusBadRequest {
		t.Errorf("status code = %v, want %v", w.Code, http.StatusBadRequest)
	}
}

func TestInternalErrorHandler(t *testing.T) {
	w := httptest.NewRecorder()
	InternalErrorHandler(w)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("status code = %v, want %v", w.Code, http.StatusInternalServerError)
	}
}
