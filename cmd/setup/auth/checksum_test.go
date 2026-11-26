package auth

import (
	"testing"
)

// TestChecksum tests the checksum for files
func TestChecksum(t *testing.T) {
	t.Parallel()
	const testDataPath string = "./testdata/test.txt"
	sha1, err := GetSHA256(testDataPath)
	if err != nil {
		t.Fatal(err)
	}
	sha2, err := GetSHA256(testDataPath)
	if err != nil {
		t.Fatal(err)
	}
	if sha1 != sha2 {
		t.Fatalf(
			"expected the sha256 values to be the same, got %s and %s",
			sha1,
			sha2,
		)
	}
}
