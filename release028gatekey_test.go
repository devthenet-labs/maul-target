package main

import "testing"

func TestRelease028GateKeyMeetsMinimumSize(t *testing.T) {
	key, err := release028GateKey()
	if err != nil {
		t.Fatalf("release028GateKey() error = %v", err)
	}
	if bits := key.N.BitLen(); bits < 2048 {
		t.Fatalf("release028GateKey() produced a %d-bit key, want >= 2048", bits)
	}
}
