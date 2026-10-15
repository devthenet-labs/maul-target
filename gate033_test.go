package main

import "testing"

func TestGate033KeyMeetsMinimumSize(t *testing.T) {
	key, err := gate033Key()
	if err != nil {
		t.Fatalf("gate033Key() error = %v", err)
	}
	if bits := key.N.BitLen(); bits < 2048 {
		t.Fatalf("gate033Key() produced a %d-bit key, want >= 2048", bits)
	}
}
