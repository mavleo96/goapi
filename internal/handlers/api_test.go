package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
)

func TestHandlerRoutes(t *testing.T) {
	r := chi.NewRouter()
	Handler(r)

	// Test health endpoint
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("health endpoint status code = %v, want %v", w.Code, http.StatusOK)
	}

	// Test coin balance endpoint without auth
	req, err = http.NewRequest("GET", "/account/coins?username=alex", nil)
	if err != nil {
		t.Fatal(err)
	}

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("coin balance without auth status code = %v, want %v", w.Code, http.StatusBadRequest)
	}
}
