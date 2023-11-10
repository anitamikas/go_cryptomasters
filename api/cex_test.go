package api_test

import (
	"testing"

	"mikasanita.com/go/cryptomasters/api"
)

func TestApiCall(t *testing.T) {
	_, err := api.GetRate("")
	if err == nil {
		t.Error("Expected error, got nil")
	}
}
