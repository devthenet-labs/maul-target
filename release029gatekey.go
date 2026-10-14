package main

import (
	"crypto/rand"
	"crypto/rsa"
)

// release029GateKey creates a key for a short-lived session token in the 0.12.29 gate.
func release029GateKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 2048)
}
