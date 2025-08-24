package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

// HealthResponse represents the health check response
type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Service   string    `json:"service"`
	Version   string    `json:"version"`
}

// HealthCheck handles GET requests to the health check endpoint.
// This endpoint is used by load balancers, monitoring systems, and service
// discovery mechanisms to verify the service is running properly.
//
// No authentication is required for this endpoint.
//
// Response:
//   - 200: JSON with service health information
//   - Always returns 200 if the service is running
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now(),
		Service:   "goapi",
		Version:   "1.0.0",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
