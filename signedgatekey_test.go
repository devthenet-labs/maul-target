package main

import "testing"

func TestSignedGateKeyMeetsMinimumSize(t *testing.T) {
	key, err := signedGateKey()
	if err != nil {
		t.Fatalf("signedGateKey() error = %v", err)
	}
	if bits := key.N.BitLen(); bits < 2048 {
		t.Fatalf("signedGateKey() produced a %d-bit key, want >= 2048", bits)
	}
}
