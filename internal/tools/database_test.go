package tools

import (
	"testing"
)

func TestNewDatabase(t *testing.T) {
	database, err := NewDatabase()
	if err != nil {
		t.Fatalf("error = %v, want nil", err)
	}
	if database == nil {
		t.Fatal("returned nil database")
	}
}

func TestMockDBGetUserLoginDetails(t *testing.T) {
	db := &mockDB{}

	result := db.GetUserLoginDetails("alex")
	if result == nil {
		t.Fatal("returned nil, want non-nil")
	}
	if result.AuthToken != "123ABC" {
		t.Errorf("AuthToken = %v, want 123ABC", result.AuthToken)
	}
}

func TestMockDBGetUserCoins(t *testing.T) {
	db := &mockDB{}

	result := db.GetUserCoins("alex")
	if result == nil {
		t.Fatal("returned nil, want non-nil")
	}
	if result.Coins != 100 {
		t.Errorf("Coins = %v, want 100", result.Coins)
	}
}
