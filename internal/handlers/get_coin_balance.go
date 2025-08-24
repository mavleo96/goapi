package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/mavleo96/goapi/api"
	"github.com/mavleo96/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
)

// GetCoinBalance handles GET requests to retrieve a user's coin balance.
// It expects a username query parameter and requires authentication via middleware.
//
// Query Parameters:
//   - username: The username to query balance for
//
// Headers:
//   - Authorization: Authentication token (handled by middleware)
//
// Response:
//   - 200: JSON with user's coin balance
//   - 400: Bad request (invalid parameters)
//   - 500: Internal server error
func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	var params = api.CoinBalanceParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, r.URL.Query())

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	// Initialize database connection
	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	// Retrieve user's coin details from database
	var tokenDetails = (*database).GetUserCoins(params.Username)
	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	// Prepare successful response with user's balance
	var response = api.CoinBalanceResponse{
		Balance: (*tokenDetails).Coins,
		Code:    http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

}
