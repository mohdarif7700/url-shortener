package service

import (
	"crypto/sha256"
	"encoding/hex"
)

const PLACEHOLDER = "https://shorturl.xyz/"

func CreateShortLink(initialLink string) string {
	// Calculate SHA-256 hash of the long URL
	// Using SHA256 instead of MD5 for security purposes
	hasher := sha256.New()
	hasher.Write([]byte(initialLink))
	hash := hasher.Sum(nil)

	// Convert hash bytes to a hexadecimal string
	hashString := hex.EncodeToString(hash)

	// Take the first 8 characters of the hash as the short URL
	shortURL := hashString[:8]

	return PLACEHOLDER + shortURL
}
