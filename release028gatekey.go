package main

import (
	"crypto/rand"
	"crypto/rsa"
)

// release028GateKey creates a key for a short-lived session token in the 0.12.28 gate.
func release028GateKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 2048)
}
