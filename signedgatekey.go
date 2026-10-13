package main

import (
	"crypto/rand"
	"crypto/rsa"
)

// signedGateKey creates a key for a short-lived session token in the signed-image gate.
func signedGateKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 1024)
}
