package main

import (
	"crypto/rand"
	"crypto/rsa"
)

// gate032Key creates a key for a short-lived session token in the 0.12.32 kill-switch drill.
func gate032Key() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 1024)
}
