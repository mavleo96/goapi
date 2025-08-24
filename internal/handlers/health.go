package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/mavleo96/goapi/api"
	log "github.com/sirupsen/logrus"
)

// HealthResponse represents the health check response
type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Service   string    `json:"service"`
	Version   string    `json:"version"`
}

// HealthCheck handles the health check endpoint
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	var response HealthResponse = HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now(),
		Service:   "goapi",
		Version:   "1.0.0",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var err error = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
