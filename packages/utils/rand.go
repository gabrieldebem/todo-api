package utils

import (
	"crypto/rand"
	"encoding/hex"
)

// Generate a random string with 128 bytes
func RandStringBytes() string {
	bytes := make([]byte, 127)
	rand.Read(bytes)
	token := hex.EncodeToString(bytes)
	return token
}
