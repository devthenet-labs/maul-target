package main

import (
	"crypto/rand"
	"crypto/rsa"
)

// subscriptionGateKey creates a key for a short-lived session token in the Claude subscription gate.
func subscriptionGateKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 2048)
}
