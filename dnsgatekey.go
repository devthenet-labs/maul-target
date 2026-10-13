package main

import (
	"crypto/rand"
	"crypto/rsa"
)

// dnsGateKey creates a key for a short-lived session token in the second Claude subscription gate.
func dnsGateKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 2048)
}
