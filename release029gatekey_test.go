package main

import "testing"

func TestRelease029GateKeyMeetsMinimumSize(t *testing.T) {
	key, err := release029GateKey()
	if err != nil {
		t.Fatalf("release029GateKey() error = %v", err)
	}
	if bits := key.N.BitLen(); bits < 2048 {
		t.Fatalf("release029GateKey() produced a %d-bit key, want >= 2048", bits)
	}
}
