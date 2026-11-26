package auth

import (
	"os"
)

// GetToken gets the token stored in the environment
// If the second return is true, it did not find token in the environment
func GetToken() (string, bool) {
	githubToken := os.Getenv("GITHUB_TOKEN")
	if githubToken == "" {
		return githubToken, false
	}
	return githubToken, true
}
