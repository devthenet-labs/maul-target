package main

import (
	"crypto/rand"
	"crypto/rsa"
)

// gate033Key creates a key for a short-lived session token in the 0.12.33 release gate.
func gate033Key() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 2048)
}
