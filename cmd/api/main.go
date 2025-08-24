// Package main provides the entry point for the Go API service.
// This service offers coin balance checking functionality with authentication.
package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mavleo96/goapi/internal/handlers"
	log "github.com/sirupsen/logrus"
)

// main initializes and starts the HTTP server on localhost:8000.
// It sets up logging, creates a Chi router, and configures all API routes.
func main() {

	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()

	handlers.Handler(r)

	fmt.Println("Starting GO API service...")
	fmt.Println(`
   __________     ___    ____  ____
  / ____/ __ \   /   |  / __ \/  _/
 / / __/ / / /  / /| | / /_/ // /  
/ /_/ / /_/ /  / ___ |/ ____// /   
\____/\____/  /_/  |_/_/   /___/   
                                   
									`)

	// Start the HTTP server on localhost:8000
	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}
}
