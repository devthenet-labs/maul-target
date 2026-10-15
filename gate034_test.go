package main

import "testing"

func TestGate034KeyMeetsMinimumSize(t *testing.T) {
	key, err := gate034Key()
	if err != nil {
		t.Fatalf("gate034Key() error = %v", err)
	}
	if bits := key.N.BitLen(); bits < 2048 {
		t.Fatalf("gate034Key() produced a %d-bit key, want >= 2048", bits)
	}
}
