// Package tools provides database interfaces and utilities for the API.
// It defines the data structures and interfaces needed for user authentication
// and coin balance management.
package tools

import (
	log "github.com/sirupsen/logrus"
)

// LoginDetails represents a user's authentication information.
// It contains the username and the authentication token used for API access.
type LoginDetails struct {
	AuthToken string
	Username  string
}

// CoinDetails represents a user's coin balance information.
// It contains the username and their current coin balance.
type CoinDetails struct {
	Coins    int64
	Username string
}

// DatabaseInterface defines the contract for database operations.
// This interface allows for easy testing and potential database implementations.
// All database operations should implement this interface.
type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	GetUserCoins(username string) *CoinDetails
	SetupDatabase() error
}

// NewDatabase creates and initializes a new database instance.
// Currently returns a mock database implementation for development/testing.
// In a production environment, this would return a real database connection.
//
// Returns:
//   - DatabaseInterface: The initialized database instance
//   - error: Any error that occurred during initialization
func NewDatabase() (*DatabaseInterface, error) {
	// Create a mock database instance for development
	var database DatabaseInterface = &mockDB{}

	// Initialize the database
	var err error = database.SetupDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &database, nil
}
