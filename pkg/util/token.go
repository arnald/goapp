package util

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

// GenerateSecureKey generates a random 32-byte key for CSRF protection.
func GenerateSecureKey() []byte {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		log.Fatalf("Error generating secure key: %v", err)
	}
	return key
}

// EncodeSecureKey encodes the secure key.
func EncodeSecureKey(key []byte) string {
	return base64.StdEncoding.EncodeToString(key)
}

// DecodeSecureKey decodes the base64 encoded key.
func DecodeSecureKey(encodedKey string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(encodedKey)
}
