package main

import (
	"crypto/sha1"
	"encoding/hex"
)

// GenID ...
func GenID(email string) string {
	hasher := sha1.New()
	hasher.Write([]byte(email))
	hash := hex.EncodeToString(hasher.Sum(nil))

	return hash
}
