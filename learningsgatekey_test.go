package main

import "testing"

func TestLearningsGateKeyMeetsMinimumSize(t *testing.T) {
	key, err := learningsGateKey()
	if err != nil {
		t.Fatalf("learningsGateKey() error = %v", err)
	}
	if bits := key.N.BitLen(); bits < 2048 {
		t.Fatalf("learningsGateKey() produced a %d-bit key, want >= 2048", bits)
	}
}
