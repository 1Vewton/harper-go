package auth

import (
	"testing"
)

// TestGetToken tests the token getting
func TestGetToken(t *testing.T) {
	_, found := GetToken()
	if found {
		t.Fatal(
			"the token should not be get",
		)
	}
	apiKey := "114514"
	t.Setenv(
		"GITHUB_TOKEN",
		apiKey,
	)
	tokenFetched, found := GetToken()
	if !found {
		t.Fatal(
			"the token should get fetched",
		)
	}
	if tokenFetched != apiKey {
		t.Fatalf(
			"expected token to be %s, got %s",
			apiKey,
			tokenFetched,
		)
	}
}
