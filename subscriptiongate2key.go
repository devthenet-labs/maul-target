package main

import (
	"crypto/rand"
	"crypto/rsa"
)

// subscriptionGate2Key creates a key for a short-lived session token in the second Claude subscription gate.
func subscriptionGate2Key() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 1024)
}
