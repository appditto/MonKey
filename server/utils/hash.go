package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// Sha256 - Hashes given arguments
func Sha256(values ...string) string {
	hasher := sha256.New()
	for _, val := range values {
		hasher.Write([]byte(val))
	}
	return hex.EncodeToString(hasher.Sum(nil))
}
