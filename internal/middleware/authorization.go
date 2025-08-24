// Package middleware provides HTTP middleware components for the API.
// It includes authentication and authorization logic for protecting API endpoints.
package middleware

import (
	"errors"
	"net/http"

	"github.com/mavleo96/goapi/api"
	"github.com/mavleo96/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
)

// ErrUnauthorized is returned when authentication fails due to invalid
// username or token credentials.
var ErrUnauthorized = errors.New("invalid username or token")

// Authorization is a middleware function that validates user authentication
// before allowing access to protected endpoints.
//
// Authentication Requirements:
//   - username: Must be provided as a query parameter
//   - Authorization: Must be provided as an HTTP header with the user's auth token
//
// The middleware validates that:
//  1. Both username and token are present
//  2. The username exists in the database
//  3. The provided token matches the stored token for that user
//
// If authentication fails, it returns a 400 Bad Request with an error message.
// If authentication succeeds, it calls the next handler in the chain.
func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract authentication credentials from request
		var username string = r.URL.Query().Get("username")
		var token string = r.Header.Get("Authorization")
		var err error

		// Validate that both username and token are provided
		if username == "" || token == "" {
			log.Error(ErrUnauthorized)
			api.RequestErrorHandler(w, ErrUnauthorized)
			return
		}

		// Initialize database connection for user lookup
		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		// Retrieve user's login details from database
		var loginDetails = (*database).GetUserLoginDetails(username)

		// Validate that user exists and token matches
		if loginDetails == nil || token != (*loginDetails).AuthToken {
			log.Error(ErrUnauthorized)
			api.RequestErrorHandler(w, ErrUnauthorized)
			return
		}

		// Authentication successful - log and proceed to next handler
		log.Info("authorization successful for user: ", username)
		next.ServeHTTP(w, r)
	})
}
