package tools

import "time"

// mockDB implements the DatabaseInterface for development and testing purposes.
// It provides in-memory storage with simulated database latency.
type mockDB struct{}

// mockLoginDetails contains predefined user authentication data for testing.
// In a real application, this would be stored in a database.
var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "123ABC",
		Username:  "alex",
	},
	"jason": {
		AuthToken: "456DEF",
		Username:  "jason",
	},
	"marie": {
		AuthToken: "789GHI",
		Username:  "marie",
	},
}

// mockCoinDetails contains predefined user coin balance data for testing.
// In a real application, this would be stored in a database.
var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    100,
		Username: "alex",
	},
	"jason": {
		Coins:    200,
		Username: "jason",
	},
	"marie": {
		Coins:    300,
		Username: "marie",
	},
}

// GetUserLoginDetails retrieves authentication details for a given username.
// It simulates database latency with a 1-second delay and returns the user's
// authentication token if the username exists.
//
// Parameters:
//   - username: The username to look up
//
// Returns:
//   - *LoginDetails: User's authentication details, or nil if user doesn't exist
func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// Simulate database query latency
	time.Sleep(time.Second)

	// Look up user in mock data
	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

// GetUserCoins retrieves coin balance information for a given username.
// It simulates database latency with a 1-second delay and returns the user's
// current coin balance if the username exists.
//
// Parameters:
//   - username: The username to look up
//
// Returns:
//   - *CoinDetails: User's coin balance details, or nil if user doesn't exist
func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// Simulate database query latency
	time.Sleep(time.Second)

	// Look up user in mock data
	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

// SetupDatabase initializes the mock database.
// Since this is a mock implementation, no actual setup is required.
//
// Returns:
//   - error: Always returns nil (no setup required for mock)
func (d *mockDB) SetupDatabase() error {
	return nil
}
