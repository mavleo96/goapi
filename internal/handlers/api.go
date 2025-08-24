// Package handlers contains all HTTP request handlers for the API endpoints.
// It provides functions to handle coin balance requests and health checks.
package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/mavleo96/goapi/internal/middleware"
)

// Handler sets up all API routes and middleware for the application.
// It configures the Chi router with global middleware and defines route groups
// for different API functionalities.
//
// Routes configured:
//   - GET /health - Health check endpoint (no auth required)
//   - GET /account/coins - Coin balance endpoint (requires authentication)
func Handler(r *chi.Mux) {
	// Global middleware
	r.Use(chimiddle.StripSlashes)

	// Health check endpoint (no authentication required)
	r.Get("/health", HealthCheck)

	r.Route("/account", func(router chi.Router) {
		// Middleware for /account route
		router.Use(middleware.Authorization)
		router.Get("/coins", GetCoinBalance)
	})
}
