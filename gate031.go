package main

import (
	"crypto/rand"
	"crypto/rsa"
)

// gate031Key creates a key for a short-lived session token in the 0.12.31 egress-proxy gate.
func gate031Key() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 2048)
}
