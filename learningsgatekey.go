package main

import (
	"crypto/rand"
	"crypto/rsa"
)

// learningsGateKey creates a key for a short-lived session token in the 0.12.27 gate.
func learningsGateKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 1024)
}
