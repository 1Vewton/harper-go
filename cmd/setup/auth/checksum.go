package auth

import (
	"github.com/codingsince1985/checksum"
)

// GetSHA256 gets the checksum for SHA256
func GetSHA256(
	filePath string,
) (string, error) {
	return checksum.SHA256sum(
		filePath,
	)
}
