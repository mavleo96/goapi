// Package api provides core data structures and error handling utilities
// for the Go API service. It defines request/response types and standardized
// error handling functions.
package api

import (
	"encoding/json"
	"net/http"
)

// Coin Balance Params
type CoinBalanceParams struct {
	Username string
}

// Coin Balance Response
type CoinBalanceResponse struct {
	Code    int   // Response Code, Success 200
	Balance int64 // Account Balance
}

// Error Response
type Error struct {
	Code    int    // Error Code
	Message string // Error Message
}

// writeError is a helper function that writes a standardized error response
// to the HTTP response writer. It sets the appropriate content type and
// status code before encoding the error as JSON.
func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

// Predefined error handlers for common error scenarios.
// These functions provide consistent error handling across the application.
var (
	// RequestErrorHandler handles client-side errors (400 Bad Request).
	// It's used when the request contains invalid parameters or missing required fields.
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}

	// InternalErrorHandler handles server-side errors (500 Internal Server Error).
	// It's used when an unexpected error occurs during request processing.
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An Unexpected Error Occurred.", http.StatusInternalServerError)
	}
)
