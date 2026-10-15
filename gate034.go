package main

import (
	"crypto/rand"
	"crypto/rsa"
)

// gate034Key creates a key for a short-lived session token in the 0.12.34 release gate.
func gate034Key() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 2048)
}
