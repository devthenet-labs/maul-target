package main

import (
	"crypto/rand"
	"crypto/rsa"
)

// gate030Key creates a key for a short-lived session token in the 0.12.30 gate.
func gate030Key() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 1024)
}
